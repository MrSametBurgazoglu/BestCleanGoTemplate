package services

import (
	"BestCleanTemplate/pkg/repository"
	"BestCleanTemplate/pkg/services/Todo"
)

type Services struct {
	TodoService *Todo.TodoService
}

func InitServices(repositories *repository.Repositories) *Services {
	services := new(Services)
	services.TodoService = &Todo.TodoService{TodoRepository: repositories.TodoRepository}
	return services
}
