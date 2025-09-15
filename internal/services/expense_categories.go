package services

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
)

type ExpenseCategoryService interface {
	CreateCategory(createCategoryDto *models.CreateExpenseCategoryDto) error
	GetCategoryByID(id int) (*models.ExpenseCategory, error)
	GetAllCategories(queryOptions *models.QueryOptions) ([]*models.ExpenseCategory, int64, error)
	DeleteCategory(id uint) error
	UpdateCategory(category *models.ExpenseCategory) error
}

type expenseCategoryService struct {
	repo repositories.ExpenseCategoryRepository
}

func NewExpenseCategoryService(_repo repositories.ExpenseCategoryRepository) ExpenseCategoryService {
	return &expenseCategoryService{
		repo: _repo,
	}
}

func (s *expenseCategoryService) CreateCategory(createCategoryDto *models.CreateExpenseCategoryDto) error {
	expenseCategory := &models.ExpenseCategory{
		Name:        createCategoryDto.Name,
		Description: createCategoryDto.Description,
	}
	return s.repo.CreateCategory(expenseCategory)
}

func (s *expenseCategoryService) GetCategoryByID(id int) (*models.ExpenseCategory, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *expenseCategoryService) GetAllCategories(queryOptions *models.QueryOptions) ([]*models.ExpenseCategory, int64, error) {
	//Query options logic
	page := queryOptions.Page
	pageSize := queryOptions.PageSize
	return s.repo.GetAllCategories(page, pageSize)
}

func (s *expenseCategoryService) UpdateCategory(updateCategoryDto *models.ExpenseCategory) error {
	expenseCategory := &models.ExpenseCategory{
		ID:          updateCategoryDto.ID,
		Name:        updateCategoryDto.Name,
		Description: updateCategoryDto.Description,
	}
	return s.repo.UpdateCategory(expenseCategory)
}

func (s *expenseCategoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}
