package services

import (
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
)

type RevenueService interface {
	CreateRevenue(createRevenueDto *models.CreateRevenueDto) error
	GetRevenueByID(id models.ID) (*models.Revenue, error)
	GetAllRevenues(queryOptions *models.QueryOptions) ([]*models.Revenue, int64, error)
	DeleteRevenue(id models.ID) error
	UpdateRevenue(id models.ID, newRevenue *models.CreateRevenueDto) error
}

type revenueService struct {
	repo repositories.RevenueRepository
}

func NewRevenueService(revenueRepo repositories.RevenueRepository) RevenueService {
	return &revenueService{
		repo: revenueRepo,
	}
}

func (s *revenueService) CreateRevenue(createRevenueDto *models.CreateRevenueDto) error {
	revenue := &models.Revenue{
		Source:      createRevenueDto.Source,
		Description: createRevenueDto.Description,
		Amount:      createRevenueDto.Amount,
	}
	return s.repo.CreateRevenue(revenue)
}

func (s *revenueService) GetRevenueByID(id models.ID) (*models.Revenue, error) {
	return s.repo.GetRevenueByID(id)
}

func (e *revenueService) GetAllRevenues(queryOptions *models.QueryOptions) ([]*models.Revenue, int64, error) {
	return e.repo.GetPaginatedRevenues(queryOptions.Page, queryOptions.PageSize)
}

func (s *revenueService) UpdateRevenue(id models.ID, newRevenue *models.CreateRevenueDto) error {
	revenue := &models.Revenue{
		ID:          id,
		Source:      newRevenue.Source,
		Description: newRevenue.Description,
		Amount:      newRevenue.Amount,
	}
	return s.repo.UpdateRevenue(revenue)
}

func (s *revenueService) DeleteRevenue(id models.ID) error {
	return s.repo.DeleteRevenue(id)
}
