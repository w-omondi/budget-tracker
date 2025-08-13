package models

import "time"

type Expense struct {
	ID          uint      `json:"id"`
	Amount      float64   `json:"amount"`
	Description string   `json:"description"`
	Date        time.Time `json:"date"`
	CategoryId    uint    `json:"category"`
}

type CreateExpenseDto struct {
	Amount      float64   `json:"amount" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	CategoryId  uint      `json:"category_id" validate:"required"`
}
