package Todo

import (
	"BestCleanTemplate/pkg/services/Todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTodoHandler struct {
	TodoService *Todo.TodoService
}

type GetTodoInput struct {
	ID int `json:"id" binding:"required"`
}

func (receiver *GetTodoHandler) GetTodo(c *gin.Context) {
	var input GetTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := receiver.TodoService.GetTodo(uint(input.ID))
	c.JSON(response.Code, response.GetJson())
}
