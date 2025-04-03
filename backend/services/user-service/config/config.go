package config

import (
	"os"
)

// Separate SMTP configuration into its own struct
type SMTPConfig struct {
	Server string
	Port   string
	User   string
	Pass   string
}

// Config holds application configuration
type Config struct {
	DBURL         string
	RedisAddr     string
	RedisPassword string
	RabbitMQURL   string
	SMTP          SMTPConfig
	APIBaseURL    string
	JWTSecret     string `mapstructure:"JWT_SECRET"`
}

// LoadConfig loads environment variables into the Config struct
func LoadConfig() Config {
	return Config{
		DBURL:         getEnv("DB_URL", "postgres://user:pass@localhost:5432/shopnest?sslmode=disable"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RabbitMQURL:   getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		SMTP: SMTPConfig{
			Server: getEnv("SMTP_SERVER", "smtp.gmail.com"),
			Port:   getEnv("SMTP_PORT", "587"),
			User:   getEnv("SMTP_USER", ""),
			Pass:   getEnv("SMTP_PASS", ""),
		},
		APIBaseURL: getEnv("API_BASE_URL", "http://localhost:8080"),
		JWTSecret:  getEnv("JWT_SECRET", "change-me-in-production"),
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
