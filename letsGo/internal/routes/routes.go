package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/DhavalSuthar-24/letsGo/internal/controllers"
	"github.com/DhavalSuthar-24/letsGo/internal/middlewares"
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

		// Auth routes
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login) // Changed from GET to POST

		// Todo routes (protected)
		todoGroup := api.Group("/todos")
		todoGroup.Use(middlewares.AuthMiddleware())
		{
			todoGroup.POST("/", todoController.CreateTodo)
			todoGroup.GET("/", todoController.GetTodos)
		}
	}

	return r
}
