package main

import (
	"github.com/w-omondi/budget-tracker.git/internal/app"
	"github.com/w-omondi/budget-tracker.git/internal/configs"
)

func main() {
	// Load environment variables
	configs.LoadEnv()
	// Initialize the application and its routes
	app.Run()
}
