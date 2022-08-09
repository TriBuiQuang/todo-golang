package portsRestFulTask

import (
	"net/http"
	"togo/internal/core/domain"
	serviceTasks "togo/internal/core/services/tasks"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
)

type TaskPort struct {
	TaskService interface {
		GetAllTasks() ([]domain.STask, int, error)
	}
}

func GetAllTasks(c *gin.Context) {
	service := &serviceTasks.TaskService{}

	tasks, totalTasks, err := service.GetAllTasks()

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

	checkTask := &task

	if checkTask.Name == "" || checkTask.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Name or user_id must not empty",
		})
		return
	}

	newTask, err := serviceTasks.CreateTask(&task)

	if err != nil {
		if err.Error() == "this user is out of the limit in order to create a new task. " {
			c.JSON(portsRestFul.PrintErrResponse(err, http.StatusBadRequest))
		} else {
			c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		}

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

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	userId := c.Param("userId")

	task := &domain.STask{ID: taskId, UserID: userId}

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
