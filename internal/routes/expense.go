package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/handlers"
	"github.com/w-omondi/budget-tracker.git/internal/repositories"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"gorm.io/gorm"
)

func ExpenseRoutes(route fiber.Router, db *gorm.DB) {
	expenseRepository := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepository)
	handler := handlers.NewExpenseHandler(expenseService)

	route.Post("/", handler.CreateExpenseHandler)
	route.Get("/:id", handler.GetExpenseByIDHandler)
	route.Get("/", handler.GetAllExpensesHandler)
	route.Put("/:id", handler.UpdateExpenseHandler)
	route.Delete("/:id", handler.DeleteExpenseHandler)
}

func ExpenseCategoriesRoute(route fiber.Router, db *gorm.DB) {
	expenseCategoriesRepo := repositories.NewExpenseCategoryRepository(db)
	expenseCategoriesService := services.NewExpenseCategoryService(expenseCategoriesRepo)
	handler := handlers.NewExpenseCategoriesHandler(expenseCategoriesService)

	route.Post("/", handler.CreateExpenseCategoriesHandler)
	route.Get("/:id", handler.GetExpenseCategoryByIDHandler)
	route.Get("/", handler.GetAllExpenseCategoriesHandler)
	route.Put("/:id", handler.UpdateExpenseCategoriesHandler)
	route.Delete("/:id", handler.DeleteExpenseCategoriesHandler)
}
