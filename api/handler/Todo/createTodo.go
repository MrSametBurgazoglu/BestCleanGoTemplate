package Todo

import (
	"BestCleanTemplate/pkg/services/Todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTodoHandler struct {
	TodoService *Todo.TodoService
}

type CreateTodoInput struct {
	TodoText string `json:"todo_text" binding:"required"`
}

func (receiver *CreateTodoHandler) CreateTodo(c *gin.Context) {
	var input CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := receiver.OTPService.Authentication(input.Code, input.ClientKey)
	c.JSON(response.Code, response.GetJson())
}
