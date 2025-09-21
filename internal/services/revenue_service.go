package services

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
)

type RevenueService interface {
	CreateRevenue(createRevenueDto *models.CreateRevenueDto) error
	GetRevenueByID(id int) (*models.Revenue, error)
	GetAllRevenues(queryOptions *models.QueryOptions) ([]*models.Revenue, int64, error)
	UpdateRevenue(id int, revenueDto *models.UpdateRevenueDto) error
	DeleteRevenue(id uint) error
}

type revenueService struct {
	repo repositories.RevenueRepository
}

func NewRevenueService(revenueRepo repositories.RevenueRepository) RevenueService {
	return &revenueService{
		repo: revenueRepo,
	}
}

func (r *revenueService) CreateRevenue(createRevenueDto *models.CreateRevenueDto) error {
	revenue := &models.Revenue{
		Amount:      createRevenueDto.Amount,
		Description: createRevenueDto.Description,
		Source:      createRevenueDto.Source,
	}
	return r.repo.CreateRevenue(revenue)
}

func (r *revenueService) GetRevenueByID(id int) (*models.Revenue, error) {
	return r.repo.GetRevenueByID(id)
}

func (r *revenueService) GetAllRevenues(queryOptions *models.QueryOptions) ([]*models.Revenue, int64, error) {
	return r.repo.GetAllRevenues(queryOptions.Page, queryOptions.PageSize)
}

func (r *revenueService) UpdateRevenue(id int, revenueDto *models.UpdateRevenueDto) error {
	revenue := &models.Revenue{
		ID:          uint(id),
		Amount:      revenueDto.Amount,
		Description: revenueDto.Description,
		Source:      revenueDto.Source,
	}
	return r.repo.UpdateRevenue(revenue)
}

func (r *revenueService) DeleteRevenue(id uint) error {
	return r.repo.DeleteRevenue(id)
}
