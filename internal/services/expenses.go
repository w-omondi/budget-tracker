package services

import (
	"time"

	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
)

type ExpenseService interface {
	CreateExpense(createExpenseDto *models.CreateExpenseDto) error
	GetExpenseByID(id int) (*models.Expense, error)
	GetAllExpenses() ([]models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id int) error
}

type expenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(expenseRepo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{
		repo: expenseRepo,
	}
}

func (e *expenseService) CreateExpense(createExpenseDto *models.CreateExpenseDto) error {
	expense := &models.Expense{
		Amount:      createExpenseDto.Amount,
		Description: createExpenseDto.Description,
		Date:        time.Now(),
		CategoryId:  createExpenseDto.CategoryId,
	}
	return e.repo.CreateExpense(expense)
}

func (e *expenseService) GetExpenseByID(id int) (*models.Expense, error) {
	return e.repo.GetExpenseByID(id)
}

func (e *expenseService) GetAllExpenses() ([]models.Expense, error) {
	return e.repo.GetAllExpenses()
}

func (e *expenseService) UpdateExpense(expense *models.Expense) error {
	return e.repo.UpdateExpense(expense)
}

func (e *expenseService) DeleteExpense(id int) error {
	return e.repo.DeleteExpense(id)
}
