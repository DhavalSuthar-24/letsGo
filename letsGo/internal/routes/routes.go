package routes

import (
	"github.com/gin-gonic/gin"


	"github.com/DhavalSuthar-24/letsGo/internal/controllers"
	"github.com/DhavalSuthar-24/letsGo/internal/services"
	"github.com/DhavalSuthar-24/letsGo/internal/repositories"
	"github.com/DhavalSuthar-24/letsGo/internal/config"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize layers
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// API routes
	api := r.Group("/api")
	{
		api.POST("/users", userController.CreateUser)
		api.GET("/users", userController.GetAllUsers)
	}

	return r
}