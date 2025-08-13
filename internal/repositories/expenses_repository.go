package repositories

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetExpenseByID(id int) (*models.Expense, error)
	GetAllExpenses() ([]models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id int) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) CreateExpense(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) GetExpenseByID(id int) (*models.Expense, error) {
	var expense models.Expense
	if err := r.db.First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) GetAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	if err := r.db.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r *expenseRepository) UpdateExpense(expense *models.Expense) error {
	return r.db.Save(expense).Error
}

func (r *expenseRepository) DeleteExpense(id int) error {
	var expense models.Expense
	if err := r.db.First(&expense, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&expense).Error
}

