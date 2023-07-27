package repository

import (
	"BestCleanTemplate/pkg/repository/Todo"
	"gorm.io/gorm"
)

type Repositories struct {
	TodoRepository *Todo.TodoRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	repositories := new(Repositories)
	repositories.TodoRepository = &Todo.TodoRepository{DB: db}
	return repositories
}
