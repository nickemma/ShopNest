package main

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/shopnest/user-service/config"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to RabbitMQ
	conn, err := amqp091.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare queue to ensure it exists
	_, err = ch.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// Set up SMTP authentication
	auth := smtp.PlainAuth(
		"",
		cfg.SMTP.User,
		cfg.SMTP.Pass,
		cfg.SMTP.Server,
	)

	// Set up graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Consume messages
	msgs, err := ch.Consume(
		"email_queue", // queue
		"",            // consumer
		false,         // auto-ack (false to manual ack)
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("Email worker started. Waiting for messages...")

	// Process messages
	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				log.Println("Message channel closed")
				return
			}
			processMessage(msg, auth, cfg)
		case <-sigs:
			log.Println("Shutting down email worker...")
			return
		}
	}
}

func processMessage(msg amqp091.Delivery, auth smtp.Auth, cfg config.Config) {
	defer func() {
		if err := msg.Ack(false); err != nil {
			log.Printf("Error acknowledging message: %v", err)
		}
	}()

	var data struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	if err := json.Unmarshal(msg.Body, &data); err != nil {
		log.Printf("Error decoding message: %v", err)
		return
	}

	// Construct verification URL using configurable base URL
	verifyURL := fmt.Sprintf("%s/verify-email?token=%s", cfg.APIBaseURL, data.Token)
	emailBody := fmt.Sprintf("Subject: Email Verification\n\nVerify your email: %s", verifyURL)

	// Send email
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", cfg.SMTP.Server, cfg.SMTP.Port),
		auth,
		cfg.SMTP.User,
		[]string{data.Email},
		[]byte(emailBody),
	)

	if err != nil {
		log.Printf("Failed to send email to %s: %v", data.Email, err)
		return
	}

	log.Printf("Successfully sent verification email to %s", data.Email)
}
