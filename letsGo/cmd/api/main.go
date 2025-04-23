package main

import (
	"github.com/DhavalSuthar-24/letsGo/internal/config"
	"github.com/DhavalSuthar-24/letsGo/internal/routes"
)

func main() {
	// Load env
	config.LoadEnv()

	// Connect DB
	config.ConnectDB()
	db := config.DB

	// Auto-migrate models
	db.AutoMigrate(&models.User{})

	// Setup routes
	r := routes.SetupRoutes(db)

	// Start server
	port := os.Getenv("SERVER_PORT")
	r.Run(":" + port)
}