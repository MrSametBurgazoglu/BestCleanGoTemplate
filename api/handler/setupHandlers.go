package handler

import (
	"BestCleanTemplate/api/handler/Todo"
	"BestCleanTemplate/pkg/services"
	"github.com/gin-gonic/gin"
)

func SetHandlers(engine *gin.Engine, apiServices *services.Services) {
	userApi := engine.Group("/user")
	{
		userApi.POST("/login", Todo.CreateTodoHandler{TodoService: apiServices.TodoService}.CreateTodo)
		userApi.POST("/send_otp", Todo.GetTodoHandler{TodoService: apiServices.TodoService}.GetTodo)
	}
}
