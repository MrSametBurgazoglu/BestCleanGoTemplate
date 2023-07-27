package Todo

import (
	"BestCleanTemplate/pkg/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB //*Database
}

func (receiver *TodoRepository) CreateTodo(todoText string) error {
	model := models.Todo{TodoText: todoText}
	if err := receiver.DB.Create(model).Error; err != nil {
		return err
	}
	//create user on database with .db
	return nil
}

func (receiver *TodoRepository) GetTodo(modelID uint) (*models.Todo, error) {
	model := models.Todo{ID: modelID}
	if err := receiver.DB.Find(model).Error; err != nil {
		return nil, err
	}
	//create user on database with .db
	return &model, nil
}
