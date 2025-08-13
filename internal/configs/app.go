package configs

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/routes"
)

func RunApp() {
	db := InitializeDatabase()

	app := fiber.New()

	routes.SetupRoutes(app,db)

	log.Fatal(app.Listen(":3000"))
}
