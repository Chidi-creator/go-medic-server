package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	Mongo_URI  string
	JWT_SECRET string
	JWT_EXPIRE string
	DB_NAME    string
}

var AppConfig *Config

// getting the env key
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func init() {
	//load env file

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = &Config{
		Port:       getEnv("PORT", "8080"),
		Mongo_URI:  getEnv("MONGO_URI", ""),
		JWT_SECRET: getEnv("JWT_SECRET", ""),
		JWT_EXPIRE: getEnv("JWT_EXPIRE", ""),
		DB_NAME:    getEnv("DB_NAME", ""),
	}

	if AppConfig.Mongo_URI == "" {
		log.Fatal("MONGO_URI is required but not set")
	}
	if AppConfig.JWT_EXPIRE == "" {
		log.Fatal("JWT_SECRET is required but not set")
	}
	if AppConfig.JWT_SECRET == "" {
		log.Fatal("JWT_SECRET is required but not set")
	}
	if AppConfig.DB_NAME == "" {
		log.Fatal("DB_NAME is required but not set")
	}

}
