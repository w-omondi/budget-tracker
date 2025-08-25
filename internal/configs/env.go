package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Environment variables loaded successfully")
}

func CheckEnvs(variables ...string) {
	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			log.Fatalf("Environment variable %s is not set", variable)
		}
	}
	log.Println("All required environment variables are set")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
