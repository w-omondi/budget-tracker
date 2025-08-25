package repositories

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type ExpenseCategoryRepository interface {
	CreateCategory(category *models.ExpenseCategory) error
	GetCategoryByID(id int) (*models.ExpenseCategory, error)
	GetAllCategories() ([]models.ExpenseCategory, error)
	UpdateCategory(category *models.ExpenseCategory) error
	DeleteCategory(id int) error
}

type expenseCategoryRepository struct {
	db *gorm.DB
}

func NewExpenseCategoryRepository(db *gorm.DB) ExpenseCategoryRepository {
	return &expenseCategoryRepository{db: db}
}

func (r *expenseCategoryRepository) CreateCategory(category *models.ExpenseCategory) error {
	return r.db.Create(category).Error
}

func (r *expenseCategoryRepository) GetCategoryByID(id int) (*models.ExpenseCategory, error) {
	var category models.ExpenseCategory
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}



func (r *expenseCategoryRepository) GetAllCategories() ([]models.ExpenseCategory, error) {
	var categories []models.ExpenseCategory
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *expenseCategoryRepository) UpdateCategory(category *models.ExpenseCategory) error {
	return r.db.Save(category).Error
}

func (r *expenseCategoryRepository) DeleteCategory(id int) error {
	var category models.ExpenseCategory
	if err := r.db.First(&category, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}
