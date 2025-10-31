package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseCategory struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Expenses    []*Expense `json:"expenses" gorm:"foreignKey:CategoryID"`
}

type CreateExpenseCategoryDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
