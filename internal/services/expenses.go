package services

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type ExpenseService struct {
	db gorm.DB // Assuming you're using GORM for database operations
}

func (e *ExpenseService) CreateExpense(expense *models.Expense) error {
	println("Creating expense:", expense)
	return e.db.Create(expense).Error
}

func (e *ExpenseService) GetExpenseByID(id int) (*models.Expense, error) {
	println("Fetching expense with ID:", id)
	// Simulate fetching from the database
	return &models.Expense{ID: id, Amount: 100.0, Description: "Sample Expense", Date: "2023-10-01", Category: "Food"}, nil
}

func (e *ExpenseService) GetAllExpenses() ([]models.Expense, error) {
	println("Fetching all expenses")
	// Simulate fetching from the database
	return []models.Expense{
		{ID: 1, Amount: 100.0, Description: "Sample Expense 1", Date: "2023-10-01", Category: "Food"},
		{ID: 2, Amount: 200.0, Description: "Sample Expense 2", Date: "2023-10-02", Category: "Transport"},
	}, nil
}

func (e *ExpenseService) UpdateExpense(expense *models.Expense) error {
	println("Updating expense:", expense)
	// Simulate updating the database
	return nil
}

func (e *ExpenseService) DeleteExpense(id int) error {
	println("Deleting expense with ID:", id)
	// Simulate deleting from the database
	return nil
}
