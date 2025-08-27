package models

import "gorm.io/gorm"

type ExpenseCategory struct {
	*gorm.Model
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Expenses    []*Expense `json:"expenses" gorm:"foreignKey:CategoryID"`
}

type CreateExpenseCategoryDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
