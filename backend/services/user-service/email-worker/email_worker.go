package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/shopnest/user-service/config"
)

// Enhanced configuration structure
type EmailConfig struct {
	FromName  string
	FromEmail string
	Retries   int
	Timeout   time.Duration
}

func main() {
	cfg := config.LoadConfig()

	// Initialize email configuration
	emailCfg := EmailConfig{
		FromName:  "Shopnest Team",
		FromEmail: cfg.SMTP.User,
		Retries:   3,
		Timeout:   10 * time.Second,
	}

	// RabbitMQ connection with retries
	conn, err := connectRabbitMQWithRetry(cfg.RabbitMQURL, 3)
	if err != nil {
		log.Fatalf("[RABBITMQ] Connection failed after retries: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[RABBITMQ] Channel creation failed: %v", err)
	}
	defer ch.Close()

	// Delete existing queue if it has wrong parameters
	_, err = ch.QueueDelete(
		"email_queue", // queue name
		false,         // ifUnused
		false,         // ifEmpty
		false,         // noWait
	)
	if err != nil {
		log.Printf("[RABBITMQ] Queue deletion warning: %v", err)
	}

	// Create fresh queue with correct parameters
	_, err = ch.QueueDeclare(
		"email_queue",
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		amqp091.Table{
			"x-message-ttl":           int32(86400000),  // 24h TTL
			"x-dead-letter-exchange":  "dead_letters",   // Add this
			"x-max-priority":          int32(10),        // Add this
		},
	)
	if err != nil {
		log.Fatalf("[RABBITMQ] Final queue creation failed: %v", err)
	}


	// SMTP authentication with validation
	if cfg.SMTP.User == "" || cfg.SMTP.Pass == "" {
		log.Fatal("[SMTP] Missing email credentials in configuration")
	}
	auth := smtp.PlainAuth("", cfg.SMTP.User, cfg.SMTP.Pass, cfg.SMTP.Server)

	// Graceful shutdown setup
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Message consumer with QOS control
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("[RABBITMQ] QOS setup failed: %v", err)
	}

	msgs, err := ch.Consume(
		"email_queue",
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("[RABBITMQ] Consumer registration failed: %v", err)
	}

	log.Println("[WORKER] Email worker started. Waiting for messages...")

	// Message processing loop
	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				log.Println("[WORKER] Message channel closed")
				return
			}
			go handleMessageWithRetry(msg, auth, cfg, emailCfg)

		case <-sigs:
			log.Println("[WORKER] Shutting down gracefully...")
			return
		}
	}
}

func connectRabbitMQWithRetry(url string, maxRetries int) (*amqp091.Connection, error) {
	var conn *amqp091.Connection
	var err error

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp091.Dial(url)
		if err == nil {
			return conn, nil
		}

		log.Printf("[RABBITMQ] Connection attempt %d/%d failed: %v", i+1, maxRetries, err)
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", maxRetries, err)
}

func handleMessageWithRetry(msg amqp091.Delivery, auth smtp.Auth, cfg config.Config, emailCfg EmailConfig) {
	defer msg.Ack(false) // Final acknowledgment

	for i := 0; i <= emailCfg.Retries; i++ {
		err := processEmailMessage(msg, auth, cfg, emailCfg)
		if err == nil {
			return
		}

		log.Printf("[RETRY] Attempt %d/%d failed: %v", i+1, emailCfg.Retries, err)
		if i < emailCfg.Retries {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	log.Printf("[ERROR] Permanent failure processing message: %s", msg.Body)
}

func processEmailMessage(msg amqp091.Delivery, auth smtp.Auth, cfg config.Config, emailCfg EmailConfig) error {
	var payload struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	if err := json.Unmarshal(msg.Body, &payload); err != nil {
		return fmt.Errorf("[JSON] Decoding failed: %w", err)
	}

	if !isValidEmail(payload.Email) {
		return fmt.Errorf("[VALIDATION] Invalid email address: %s", payload.Email)
	}

	// Email content generation
	verifyURL := fmt.Sprintf("%s/verify-email?token=%s", cfg.APIBaseURL, payload.Token)
	emailContent, err := createVerificationEmail(emailCfg, payload.Email, verifyURL)
	if err != nil {
		return fmt.Errorf("[EMAIL] Content creation failed: %w", err)
	}

	// SMTP connection with TLS
	client, err := smtp.Dial(fmt.Sprintf("%s:%s", cfg.SMTP.Server, cfg.SMTP.Port))
	if err != nil {
		return fmt.Errorf("[SMTP] Connection failed: %w", err)
	}
	defer client.Close()

	// STARTTLS negotiation
	if ok, _ := client.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{
			ServerName:         cfg.SMTP.Server,
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		}

		if err = client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("[SMTP] TLS handshake failed: %w", err)
		}
	}

	// Authentication
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("[SMTP] Authentication failed: %w", err)
	}

	// Set sender and recipient
	if err = client.Mail(emailCfg.FromEmail); err != nil {
		return fmt.Errorf("[SMTP] Sender setup failed: %w", err)
	}
	if err = client.Rcpt(payload.Email); err != nil {
		return fmt.Errorf("[SMTP] Recipient setup failed: %w", err)
	}

	// Send email data
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("[SMTP] Data command failed: %w", err)
	}
	defer w.Close()

	if _, err = w.Write(emailContent); err != nil {
		return fmt.Errorf("[SMTP] Writing content failed: %w", err)
	}

	log.Printf("[SUCCESS] Verification email sent to %s", payload.Email)
	return nil
}

func createVerificationEmail(cfg EmailConfig, toEmail string, verifyURL string) ([]byte, error) {
	var buffer bytes.Buffer

	// Email headers
	buffer.WriteString(fmt.Sprintf("From: \"%s\" <%s>\r\n", cfg.FromName, cfg.FromEmail))
	buffer.WriteString(fmt.Sprintf("To: %s\r\n", toEmail))
	buffer.WriteString("Subject: Verify Your Shopnest Account\r\n")
	buffer.WriteString("MIME-Version: 1.0\r\n")
	buffer.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")

	// HTML template
	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<style>
			.container { max-width: 600px; margin: 20px auto; padding: 20px; }
			.button { background-color: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Email Verification</h2>
			<p>Please click the button below to verify your email address:</p>
			<a href="%s" class="button">Verify Email</a>
			<p>If you didn't request this, please ignore this email.</p>
		</div>
	</body>
	</html>
	`, verifyURL)

	buffer.WriteString(htmlContent)
	return buffer.Bytes(), nil
}

func isValidEmail(email string) bool {
	// Simple regex-based validation
	return len(email) > 3 && bytes.Contains([]byte(email), []byte("@"))
}