package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration holds all environment variables
type Configuration struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBCharset     string
	DBParseTime   string
	DBLoc         string
	ServerPort    string
	GinMode       string
	JWTSecret     string
	JWTExpiration string
}

var Config Configuration

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	// Try to load .env file, but continue if it doesn't exist
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default or system environment variables")
	}

	// Load configuration with default values
	Config = Configuration{
		DBHost:        getEnv("DB_HOST", "127.0.0.1"),
		DBPort:        getEnv("DB_PORT", "3308"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "passroot"),
		DBName:        getEnv("DB_NAME", "go_demo"),
		DBCharset:     getEnv("DB_CHARSET", "utf8mb4"),
		DBParseTime:   getEnv("DB_PARSE_TIME", "True"),
		DBLoc:         getEnv("DB_LOC", "Local"),
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		GinMode:       getEnv("GIN_MODE", "debug"),
		JWTSecret:     getEnv("JWT_SECRET", "default_secret_key"),
		JWTExpiration: getEnv("JWT_EXPIRATION", "24h"),
	}
}

// GetDSN returns the database connection string
func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		Config.DBUser,
		Config.DBPassword,
		Config.DBHost,
		Config.DBPort,
		Config.DBName,
		Config.DBCharset,
		Config.DBParseTime,
		Config.DBLoc,
	)
}

// GetRootDSN returns a DSN without database name for creating the database
func GetRootDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=%s&loc=%s",
		Config.DBUser,
		Config.DBPassword,
		Config.DBHost,
		Config.DBPort,
		Config.DBCharset,
		Config.DBParseTime,
		Config.DBLoc,
	)
}

// Helper function to get environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
