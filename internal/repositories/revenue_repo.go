package repositories

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type RevenueRepository interface {
	CreateRevenue(revenue *models.Revenue) error
	GetRevenueByID(id models.ID) (*models.Revenue, error)
	GetPaginatedRevenues(page, pageSize int) ([]*models.Revenue, int64, error)
	UpdateRevenue(revenue *models.Revenue) error
	DeleteRevenue(id models.ID) error
}

type revenueRepository struct {
	db *gorm.DB
}

func NewRevenueRepository(db *gorm.DB) RevenueRepository {
	return &revenueRepository{db: db}
}

func (r *revenueRepository) CreateRevenue(revenue *models.Revenue) error {
	return r.db.Create(revenue).Error
}

func (r *revenueRepository) GetRevenueByID(id models.ID) (*models.Revenue, error) {
	var revenue models.Revenue
	if err := r.db.First(&revenue, id).Error; err != nil {
		return nil, err
	}
	return &revenue, nil
}

func (r *revenueRepository) GetPaginatedRevenues(page, pageSize int) ([]*models.Revenue, int64, error) {
	var revenues []*models.Revenue
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Offset(offset).Limit(pageSize).Find(&revenues).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.Revenue{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return revenues, total, nil
}

func (r *revenueRepository) UpdateRevenue(revenue *models.Revenue) error {
	return r.db.Model(&models.Revenue{}).Where("id=?", revenue.ID).Updates(revenue).Error
}

func (r *revenueRepository) DeleteRevenue(id models.ID) error {
	var revenue models.Revenue
	if err := r.db.First(&revenue, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&revenue).Error
}
