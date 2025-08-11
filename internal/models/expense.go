package models

type Expense struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
}

type CreateExpenseDto struct {
	Amount      float64 `json:"amount" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Date        string  `json:"date" validate:"required"`
	Category    string  `json:"category" validate:"required"`
}