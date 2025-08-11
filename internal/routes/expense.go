package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/handlers"
)

func ExpenseRoutes(route fiber.Router) {
	route.Post("/", handlers.CreateExpenseHandler)
	route.Get("/:id", handlers.GetExpenseByIDHandler)
	route.Get("/", handlers.GetAllExpensesHandler)
	route.Put("/:id", handlers.UpdateExpenseHandler)
	route.Delete("/:id", handlers.DeleteExpenseHandler)
}
