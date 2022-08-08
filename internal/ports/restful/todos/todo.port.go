package portsRestFulTodo

import (
	"net/http"
	"togo/internal/core/domain"
	serviceTodo "togo/internal/core/services/todos"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	todos, err := serviceTodo.GetAllTodos()

	if err != nil {

		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})
}

func CreateTodo(c *gin.Context) {
	var todo domain.STodo
	c.BindJSON(&todo)

	newTodo, err := serviceTodo.CreateTodo(todo)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
		"data":    newTodo,
	})
}

func GetSingleTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &domain.STodo{ID: todoId}

	err := serviceTodo.GetSingleTodo(todo)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data":    todo,
	})

}

func EditTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	var todo domain.STodo
	c.BindJSON(&todo)
	// completed := todo.Completed

	err := serviceTodo.EditTodo(todoId, todo.Completed)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Todo Edited Successfully",
	})

}

func DeleteTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &domain.STodo{ID: todoId}

	err := serviceTodo.DeleteTodo(todo)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
}
