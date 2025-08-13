package configs

import (
	"fmt"

	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	CheckEnvs(
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_PORT",
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	if err := db.AutoMigrate(&models.Expense{}); err != nil {
		panic("failed to migrate database")
	}

	return db
}
