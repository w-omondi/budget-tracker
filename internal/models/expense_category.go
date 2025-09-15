package models

import (
	"time"

	"gorm.io/gorm"
)

type ExpenseCategory struct {
	// *gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
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
