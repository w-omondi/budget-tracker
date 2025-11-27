package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/handlers"
	"github.com/w-omondi/budget-tracker.git/internal/middlewares"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"gorm.io/gorm"
)

func DashboardRoutes(route fiber.Router, db *gorm.DB) {
	expenseRepository := repositories.NewExpenseRepository(db)
	revenueRepository := repositories.NewRevenueRepository(db)
	categoriesRepository := repositories.NewExpenseCategoryRepository(db)

	expenseService := services.NewExpenseService(expenseRepository)
	revenueService := services.NewRevenueService(revenueRepository)
	categoryService := services.NewExpenseCategoryService(categoriesRepository)

	dashboardHandler := handlers.NewDashbordHandler(
		expenseService,
		revenueService,
		categoryService,
	)

	route.Get("/", middlewares.AuthenticateUser(), dashboardHandler.GetDashbordSummary)
}
