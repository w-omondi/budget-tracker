package repositories

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetExpenseByID(id models.ID) (*models.Expense, error)
	GetAllExpenses() ([]models.Expense, error)
	GetPaginatedExpenses(page, pageSize int) ([]*models.Expense, int64, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id models.ID) error
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

func (r *expenseRepository) GetExpenseByID(id models.ID) (*models.Expense, error) {
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

func (r *expenseRepository) GetPaginatedExpenses(page, pageSize int) ([]*models.Expense, int64, error) {
	var expenses []*models.Expense
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Preload("Category").Offset(offset).Limit(pageSize).Find(&expenses).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.Expense{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return expenses, total, nil
}

func (r *expenseRepository) UpdateExpense(expense *models.Expense) error {
	return r.db.Model(&models.Expense{}).Where("id=?", expense.ID).Updates(expense).Error
}

func (r *expenseRepository) DeleteExpense(id models.ID) error {
	var expense models.Expense
	if err := r.db.First(&expense, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&expense).Error
}
