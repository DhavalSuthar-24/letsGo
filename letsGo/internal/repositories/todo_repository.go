package repositories

import (
	"gorm.io/gorm"

	"github.com/DhavalSuthar-24/letsGo/internal/models"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}
func (r *TodoRepository) GetByUserID(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where("user_id =?", userID).Find(&todos).Error
	return todos, err
}
