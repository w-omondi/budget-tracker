package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expense struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Amount      float64          `json:"amount"`
	Description string           `json:"description"`
	CategoryID  uuid.UUID        `json:"category_id" validate:"required"`
	Category    *ExpenseCategory `json:"category" gorm:"foreignKey:CategoryID"`
	Priority    string           `json:"priority"`
}

type CreateExpenseDto struct {
	Amount      float64   `json:"amount" validate:"required,min=1"`
	Description string    `json:"description" validate:"required"`
	CategoryId  uuid.UUID `json:"category_id" validate:"required"`
	Priority    string    `json:"priority"`
}
