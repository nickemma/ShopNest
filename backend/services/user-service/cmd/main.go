package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/shopnest/user-service/api/handler"
	"github.com/shopnest/user-service/api/routes"
	"github.com/shopnest/user-service/config"
	"github.com/shopnest/user-service/internal/application"
	"github.com/shopnest/user-service/internal/repository"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db, err := pgx.Connect(context.Background(), cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
	})

	// Connect to RabbitMQ
	conn, err := amqp091.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatal(err)
	}
	// close the connection when there is no queue
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("email_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository, service, and handler
	repo := repository.NewPostgresUserRepository(db)
	service := application.NewUserService(
		repo,
		redisClient,
		ch,
		config.SMTPConfig{ // Match the struct type exactly
			Server: cfg.SMTP.Server,
			Port:   cfg.SMTP.Port,
			User:   cfg.SMTP.User,
			Pass:   cfg.SMTP.Pass,
		},
		cfg.JWTSecret,
	)

	handler := handler.NewUserHandler(service)

	// Setup and run server
	r := routes.SetupRouter(handler, cfg)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
