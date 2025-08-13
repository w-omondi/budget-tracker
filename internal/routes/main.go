package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to the Budget Tracker API!")
	})

	expenseRoutes := app.Group("/expenses")
	ExpenseRoutes(expenseRoutes, db)
}
