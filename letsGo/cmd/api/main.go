package main

import (
	"os"

	"github.com/DhavalSuthar-24/letsGo/internal/config"
	"github.com/DhavalSuthar-24/letsGo/internal/routes"
		"github.com/DhavalSuthar-24/letsGo/internal/models"
)

func main() {

	config.LoadEnv()


	config.ConnectDB()
	db := config.DB

	// Auto-migrate models
	db.AutoMigrate(&models.User{},&models.Todo{})

	// Setup routes
	r := routes.SetupRoutes(db)

	// Start server
	port := os.Getenv("SERVER_PORT")
	r.Run(":" + port)
}