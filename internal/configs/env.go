package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        // Not fatal in production. Use environment variables from runtime instead.
        log.Println("No .env loaded (this is okay in production)")
    } else {
        log.Println("Environment variables loaded from .env")
    }
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
