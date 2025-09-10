package models

import "gorm.io/gorm"

type ExpenseCategory struct {
	// *gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey"`
	UpdatedAt   string         `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedAt   string         `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description"`
	Expenses    []*Expense     `json:"expenses" gorm:"foreignKey:CategoryID"`
}

type CreateExpenseCategoryDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
