package repositories

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"gorm.io/gorm"
)

type RevenueRepository interface {
	CreateRevenue(revenue *models.Revenue) error
	GetRevenueByID(id int) (*models.Revenue, error)
	GetAllRevenues(page, pageSize int) ([]*models.Revenue, int64, error)
	UpdateRevenue(revenue *models.Revenue) error
	DeleteRevenue(id uint) error
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

func (r *revenueRepository) GetRevenueByID(id int) (*models.Revenue, error) {
	var revenue models.Revenue
	if err := r.db.First(&revenue, id).Error; err != nil {
		return nil, err
	}
	return &revenue, nil
}

func (r *revenueRepository) GetAllRevenues(page, pageSize int) ([]*models.Revenue, int64, error) {
	var revenues []*models.Revenue
	var totalRecords int64

	r.db.Model(&models.Revenue{}).Count(&totalRecords)

	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Find(&revenues).Error; err != nil {
		return nil, 0, err
	}
	return revenues, totalRecords, nil
}

func (r *revenueRepository) UpdateRevenue(revenue *models.Revenue) error {
	return r.db.Model(&models.Revenue{}).Where("id = ?", revenue.ID).Updates(revenue).Error
}

func (r *revenueRepository) DeleteRevenue(id uint) error {
	return r.db.Delete(&models.Revenue{}, id).Error
}
