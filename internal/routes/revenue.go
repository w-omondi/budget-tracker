package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/handlers"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"gorm.io/gorm"
)

func RevenueRoutes(route fiber.Router, db *gorm.DB) {
	revenueRepository := repositories.NewRevenueRepository(db)
	revenueService := services.NewRevenueService(revenueRepository)
	handler := handlers.NewRevenueHandler(revenueService)

	route.Post("/", handler.CreateRevenue)
	route.Get("/:id", handler.GetRevenueByID)
	route.Get("/", handler.GetAllRevenues)
	route.Put("/:id", handler.UpdateRevenue)
	route.Delete("/:id", handler.DeleteRevenue)
}
