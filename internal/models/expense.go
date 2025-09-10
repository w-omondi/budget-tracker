package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	// *gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	UpdatedAt string         `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedAt string         `json:"created_at" gorm:"autoCreateTime"`
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
