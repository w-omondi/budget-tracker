package configs

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/routes"
)

func RunApp() {
	app := fiber.New()

	routes.DefaultRoutes(app)

	expenseRoutes := app.Group("/expenses")
	routes.ExpenseRoutes(expenseRoutes)

	log.Fatal(app.Listen(":3000"))
}
