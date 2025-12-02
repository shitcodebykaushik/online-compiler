package configs

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	GinMode           string
	Judge0URL         string
	Judge0Timeout     int
	RedisURL          string
	RedisPassword     string
	RedisDB           int
	DatabasePath      string
	RateLimitRequests int
	RateLimitWindow   int
	AllowedOrigins    []string
}

var AppConfig *Config

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		Port:              getEnv("PORT", "8080"),
		GinMode:           getEnv("GIN_MODE", "debug"),
		Judge0URL:         getEnv("JUDGE0_URL", "http://localhost:2358"),
		Judge0Timeout:     getEnvAsInt("JUDGE0_TIMEOUT", 10),
		RedisURL:          getEnv("REDIS_URL", "localhost:6379"),
		RedisPassword:     getEnv("REDIS_PASSWORD", ""),
		RedisDB:           getEnvAsInt("REDIS_DB", 0),
		DatabasePath:      getEnv("DATABASE_PATH", "./data/compiler.db"),
		RateLimitRequests: getEnvAsInt("RATE_LIMIT_REQUESTS", 30),
		RateLimitWindow:   getEnvAsInt("RATE_LIMIT_WINDOW", 900),
		AllowedOrigins:    getEnvAsSlice("ALLOWED_ORIGINS", []string{"http://localhost:5173"}),
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsSlice(key string, defaultVal []string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultVal
	}
	return strings.Split(valueStr, ",")
}
