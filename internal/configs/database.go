package configs

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	if err := db.AutoMigrate(&models.Expense{}); err != nil {
		panic("failed to migrate database")
	}

	return db
}
