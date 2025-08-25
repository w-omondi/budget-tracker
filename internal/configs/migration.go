package configs

import (
	"log"

	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type MigrationsManager struct {
	db *gorm.DB
}

func NewMigrationManager(_db *gorm.DB) *MigrationsManager {
	return &MigrationsManager{
		db: _db,
	}
}

func (m *MigrationsManager) Run() error {
	models := []any{
		&models.Expense{},
		&models.ExpenseCategory{},
	}

	if err := m.db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
