package routes

import "github.com/gofiber/fiber/v2"

func DefaultRoutes(route *fiber.App) {
	route.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to the Budget Tracker API!")
	})
	route.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("API is running smoothly!")
	})
}
