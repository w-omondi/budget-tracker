package configs

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/middlewares"
	"github.com/w-omondi/budget-tracker.git/internal/routes"
)

func RunApp() {
	db := InitializeDatabase()
	
	app := fiber.New()
	middlewares.SetupCors(app)

	appRouter := routes.NewAppRouter(app, db)
	appRouter.CreateRouter()

	log.Fatal(app.Listen(":3001"))
}
