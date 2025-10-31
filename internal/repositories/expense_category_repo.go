package repositories

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type ExpenseCategoryRepository interface {
	CreateCategory(category *models.ExpenseCategory) error
	GetCategoryByID(id models.ID) (*models.ExpenseCategory, error)
	GetAllCategories(page, pageSize int) ([]*models.ExpenseCategory, int64, error)
	UpdateCategory(category *models.ExpenseCategory) error
	DeleteCategory(id models.ID) error
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

func (r *expenseCategoryRepository) GetCategoryByID(id models.ID) (*models.ExpenseCategory, error) {
	var category models.ExpenseCategory
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *expenseCategoryRepository) GetAllCategories(page, pageSize int) ([]*models.ExpenseCategory, int64, error) {
	var categories []*models.ExpenseCategory
	var total int64
	offset := (page - 1) * pageSize

	if err := r.db.Find(&categories).Offset(offset).Limit(pageSize).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.ExpenseCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *expenseCategoryRepository) UpdateCategory(category *models.ExpenseCategory) error {
	return r.db.Model(&models.ExpenseCategory{}).Where("id=?", category.ID).Updates(category).Error
}

func (r *expenseCategoryRepository) DeleteCategory(id models.ID) error {
	category := &models.ExpenseCategory{
		ID: id,
	}
	if err := r.db.First(category).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}
