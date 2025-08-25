package models

import "gorm.io/gorm"

type ExpenseCategory struct {
	*gorm.Model
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CreateExpenseCategoryDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
