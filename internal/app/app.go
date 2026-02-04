package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/configs"
	"github.com/w-omondi/budget-tracker.git/internal/middlewares"
	"github.com/w-omondi/budget-tracker.git/internal/routes"
)

func Run() {
	dbManager := configs.NewDbManager()
	db := dbManager.InitializeDatabase()
	dbManager.EnableUUIDExtension()
	dbManager.RunMigration()

	app := fiber.New()
	middlewares.SetupCors(app)

	appRouter := routes.NewAppRouter(app, db)
	appRouter.CreateRouter()

	log.Fatal(app.Listen(":5000"))
}
