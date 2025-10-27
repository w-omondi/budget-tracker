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

	route.Post("/", handler.CreateRevenueHandler)
	route.Get("/:id", handler.GetRevenueByIDHandler)
	route.Get("/", handler.GetAllRevenuesHandler)
	route.Put("/:id", handler.UpdateRevenueHandler)
	route.Delete("/:id", handler.DeleteRevenueHandler)
}
