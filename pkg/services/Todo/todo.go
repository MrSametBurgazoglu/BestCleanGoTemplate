package Todo

import (
	"BestCleanTemplate/api/responses"
	"BestCleanTemplate/pkg/repository/Todo"
)

type TodoService struct {
	TodoRepository *Todo.TodoRepository
}

func (receiver *TodoService) CreateTodo(todoText string) (baseResponse *responses.BaseResponse) {
	baseResponse = new(responses.BaseResponse)
	err := receiver.TodoRepository.CreateTodo(todoText)
	if err != nil {
		return baseResponse.CreateErrorResponse(err)
	}
	return baseResponse.CreateSuccessResponse("otp is successfully send")
}

func (receiver *TodoService) GetTodo(modelID uint) (baseResponse *responses.BaseResponse) {
	baseResponse = new(responses.BaseResponse)
	todoModel, err := receiver.TodoRepository.GetTodo(modelID)
	if err != nil {
		return baseResponse.CreateErrorResponse(err)
	}
	baseResponse.SetData("todo", todoModel)
	return baseResponse.CreateSuccessResponse("otp is successfully send")
}
