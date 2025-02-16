package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	MongoURI       string // MongoDB connection URI
	Port           string // Server port
	RequestTimeout time.Duration
}

// LoadConfig loads configuration from .env
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Parse REQUEST_TIMEOUT from .env
	requestTimeoutStr := os.Getenv("REQUEST_TIMEOUT")
	requestTimeout, err := time.ParseDuration(requestTimeoutStr)
	if err != nil {
		log.Println("Invalid REQUEST_TIMEOUT value, using default 5s")
		requestTimeout = 5 * time.Second
	}

	return &Config{
		MongoURI:       mongoURI,
		Port:           port,
		RequestTimeout: requestTimeout,
	}
}
