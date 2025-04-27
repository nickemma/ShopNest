package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/shopnest/user-service/api/handler"
	"github.com/shopnest/user-service/api/routes"
	"github.com/shopnest/user-service/config"
	"github.com/shopnest/user-service/internal/application"
	"github.com/shopnest/user-service/internal/repository"
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
		log.Fatal(err)
	}

	// Initialize repository, service, and handler
	customerRepo := repository.NewPostgresCustomerRepository(db)
	authRepo := repository.NewPostgresAuthRepository(db)
	managerRepo := repository.NewPostgresManagerRepository(db)

	authService := application.NewAuthService(
		authRepo,
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
	authHandler := handler.NewAuthHandler(authService)

	customerService := application.NewCustomerService(
		customerRepo,
		authRepo,
	)
	userHandler := handler.NewCustomerHandler(customerService)

	managerService := application.NewManagerService(
		managerRepo,
		authRepo,
	)

	managerHandler := handler.NewManagerHandler(managerService)

	r := gin.Default()
	r.Use(gin.Logger()) // Enable Gin's built-in logger

	api := r.Group("/api/v1")
	routes.RegisterAuthRoutes(api.Group("/auth"), authHandler, cfg)
	routes.RegisterCustomerRoutes(api.Group("/customers"), userHandler, cfg)
	routes.RegisterManagerRoutes(api.Group("/managers"), managerHandler, cfg)

	// Setup and run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
