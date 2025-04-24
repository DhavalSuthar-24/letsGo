package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/DhavalSuthar-24/letsGo/internal/controllers"
	"github.com/DhavalSuthar-24/letsGo/internal/middleware" // Fixed import path
	"github.com/DhavalSuthar-24/letsGo/internal/repositories"
	"github.com/DhavalSuthar-24/letsGo/internal/services"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize user/auth layers
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	// Initialize todo layers
	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	// API routes
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.String(200, "lets fuc*ing go with Go")
		})

		// Auth routes (unprotected)
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)

		// Todo routes (protected)
		todoGroup := api.Group("/todos")
		todoGroup.Use(middleware.AuthMiddleware()) // Proper middleware usage
		{
			todoGroup.POST("/", todoController.CreateTodo)
			todoGroup.GET("/", todoController.GetTodos)
		}
	}

	return r
}
