package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	// *gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Amount      float64          `json:"amount"`
	Description string           `json:"description"`
	CategoryID  uint             `json:"category_id" validate:"required"`
	Category    *ExpenseCategory `json:"category" gorm:"foreignKey:CategoryID"`
}

type CreateExpenseDto struct {
	Amount      float64 `json:"amount" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CategoryId  uint    `json:"category_id" validate:"required"`
}
