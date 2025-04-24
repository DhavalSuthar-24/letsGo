package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DhavalSuthar-24/letsGo/internal/models"
	"github.com/DhavalSuthar-24/letsGo/internal/services"
)

type TodoController struct {
	TodoService *services.TodoService
}

func NewTodoController(TodoService *services.TodoService) *TodoController {
	return &TodoController{TodoService: TodoService}
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get userID from middleware
	userID, _ := ctx.Get("userID")
	todo.UserID = userID.(uint)

	if err := c.TodoService.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	todos, err := c.TodoService.GetTodosByUser(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}
