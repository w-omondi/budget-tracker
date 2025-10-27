package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppRouter interface {
	CreateRouter()
}

type appRouter struct {
	app *fiber.App
	db  *gorm.DB
}

func NewAppRouter(app *fiber.App, db *gorm.DB) AppRouter {
	return &appRouter{
		app: app,
		db:  db,
	}
}

func (r *appRouter) CreateRouter() {
	app := r.app
	db := r.db

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to the Budget Tracker API!")
	})

	api := r.app.Group("/api/v1")

	revenueRoutes := api.Group("/revenue-sources")
	RevenueRoutes(revenueRoutes, db)

	expenseRoutes := api.Group("/expenses")
	ExpenseRoutes(expenseRoutes, db)

	expenseCategoriesRoutes := api.Group("/expense-categories")
	ExpenseCategoriesRoute(expenseCategoriesRoutes, db)
}
