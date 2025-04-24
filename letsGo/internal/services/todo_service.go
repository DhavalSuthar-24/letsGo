package services

import (
	"github.com/DhavalSuthar-24/letsGo/internal/models"
	"github.com/DhavalSuthar-24/letsGo/internal/repositories"
)

type TodoService struct {
	todoRepo *repositories.TodoRepository
}

func NewTodoService(todoRepo *repositories.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return s.todoRepo.Create(todo)
}
func (s *TodoService) GetTodosByUser(userID uint) ([]models.Todo, error) {
	return s.todoRepo.GetByUserID(userID)
}
