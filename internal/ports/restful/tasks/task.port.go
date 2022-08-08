package portsRestFulTask

import (
	"net/http"
	"togo/internal/core/domain"
	serviceTasks "togo/internal/core/services/tasks"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks, totalTasks, err := serviceTasks.GetAllTasks()

	if err != nil {

		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"message":    "All Tasks",
		"total_task": totalTasks,
		"data":       tasks,
	})
}

func CreateTask(c *gin.Context) {
	var task domain.STask
	c.BindJSON(&task)

	newTask, err := serviceTasks.CreateTask(task)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Task created Successfully",
		"data":    newTask,
	})
}

func GetSingleTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task := &domain.STask{ID: taskId}

	err := serviceTasks.GetSingleTask(task)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Task",
		"data":    task,
	})

}

func EditTask(c *gin.Context) {
	taskId := c.Param("taskId")
	var task domain.STask
	c.BindJSON(&task)
	// completed := task.Completed

	err := serviceTasks.EditTask(taskId, task)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Task Edited Successfully",
	})

}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task := &domain.STask{ID: taskId}

	err := serviceTasks.DeleteTask(task)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Task deleted successfully",
	})
}
