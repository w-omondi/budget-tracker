package configs

import (
	"fmt"
	"log"

	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager interface {
	InitializeDatabase() *gorm.DB
	EnableUUIDExtension()
	RunMigration() error
}

type dbManager struct {
	db *gorm.DB
}

func NewDbManager() DBManager {
	return &dbManager{}
}

func (dm *dbManager) InitializeDatabase() *gorm.DB {
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

	dm.db = db

	return db
}

func (dm *dbManager) EnableUUIDExtension() {
	sql := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	if err := dm.db.Exec(sql).Error; err != nil {
		log.Fatalf("❌ Failed to enable uuid-ossp extension: %v", err)
	}
	log.Println("✅ uuid-ossp extension is active")
}

func (dm *dbManager) RunMigration() error {
	models := []any{
		&models.Expense{},
		&models.ExpenseCategory{},
		&models.Revenue{},
	}

	if err := dm.db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
