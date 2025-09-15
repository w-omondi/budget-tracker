package services

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
)

type ExpenseService interface {
	CreateExpense(createExpenseDto *models.CreateExpenseDto) error
	GetExpenseByID(id int) (*models.Expense, error)
	GetAllExpenses(queryOptions *models.QueryOptions) ([]*models.Expense, int64, error)
	UpdateExpense(id int, expenseDto *models.CreateExpenseDto) error
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
		CategoryID:  createExpenseDto.CategoryId,
	}
	return e.repo.CreateExpense(expense)
}

func (e *expenseService) GetExpenseByID(id int) (*models.Expense, error) {
	return e.repo.GetExpenseByID(id)
}

func (e *expenseService) GetAllExpenses(queryOptions *models.QueryOptions) ([]*models.Expense, int64, error) {
	return e.repo.GetPaginatedExpenses(queryOptions.Page, queryOptions.PageSize)
}

func (e *expenseService) UpdateExpense(id int, expenseDto *models.CreateExpenseDto) error {
	expense := &models.Expense{
		ID:          uint(id),
		Amount:      expenseDto.Amount,
		Description: expenseDto.Description,
		CategoryID:  expenseDto.CategoryId,
	}
	return e.repo.UpdateExpense(expense)
}

func (e *expenseService) DeleteExpense(id int) error {
	return e.repo.DeleteExpense(id)
}
