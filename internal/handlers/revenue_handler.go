package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/services"
)

type RevenueHandler interface {
	CreateRevenue(ctx *fiber.Ctx) error
	GetRevenueByID(ctx *fiber.Ctx) error
	GetAllRevenues(ctx *fiber.Ctx) error
	UpdateRevenue(ctx *fiber.Ctx) error
	DeleteRevenue(ctx *fiber.Ctx) error
}

type revenueHandler struct {
	revenueService services.RevenueService
}

func NewRevenueHandler(service services.RevenueService) RevenueHandler {
	return &revenueHandler{
		service: service,
	}
}

func (h *revenueHandler) CreateRevenue(ctx *fiber.Ctx) error {
	
}