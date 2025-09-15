package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Mongo_URI  string
	JWT_SECRET string
	JWT_EXPIRE string
}


//getting the env key 
func getEnv(key, fallback string)string{
	if value, exists := os.LookupEnv(key); exists{
		return  value
	}
	return fallback
}



func LoadConfig() *Config {
	//load env file

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		Port: getEnv("PORT", "8080"),
		Mongo_URI: getEnv("MONGO_URI", ""),
		JWT_SECRET: getEnv("JWT_SECRET", ""),
		JWT_EXPIRE: getEnv("JWT_EXPIRE", ""),


	}

	return cfg
	
}


