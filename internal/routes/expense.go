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
