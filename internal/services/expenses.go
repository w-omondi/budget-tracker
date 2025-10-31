package services

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
)

type ExpenseService interface {
	CreateExpense(createExpenseDto *models.CreateExpenseDto) error
	GetExpenseByID(id models.ID) (*models.Expense, error)
	GetAllExpenses(queryOptions *models.QueryOptions) ([]*models.Expense, int64, error)
	UpdateExpense(id models.ID, expenseDto *models.CreateExpenseDto) error
	DeleteExpense(id models.ID) error
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
		Priority:    createExpenseDto.Priority,
	}
	return e.repo.CreateExpense(expense)
}

func (e *expenseService) GetExpenseByID(id models.ID) (*models.Expense, error) {
	return e.repo.GetExpenseByID(id)
}

func (e *expenseService) GetAllExpenses(queryOptions *models.QueryOptions) ([]*models.Expense, int64, error) {
	return e.repo.GetPaginatedExpenses(queryOptions.Page, queryOptions.PageSize)
}

func (e *expenseService) UpdateExpense(id models.ID, expenseDto *models.CreateExpenseDto) error {
	expense := &models.Expense{
		ID:          id,
		Amount:      expenseDto.Amount,
		Description: expenseDto.Description,
		CategoryID:  expenseDto.CategoryId,
		Priority:    expenseDto.Priority,
	}
	return e.repo.UpdateExpense(expense)
}

func (e *expenseService) DeleteExpense(id models.ID) error {
	return e.repo.DeleteExpense(id)
}
