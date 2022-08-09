package portsRestFulTask

import (
	"net/http"
	"togo/internal/core/domain"
	serviceTasks "togo/internal/core/services/tasks"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
)

type STaskPort struct {
	TaskService interface {
		GetAllTasks(c *gin.Context)
		CreateTask(c *gin.Context)
		GetSingleTask(c *gin.Context)
		DeleteTask(c *gin.Context)
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
	service := &serviceTasks.TaskService{}
	c.BindJSON(service.Task)

	if portsRestFul.ValidateCreateTask(c, service.Task) {
		return
	}

	newTask, err := service.CreateTask()
	// newTask, err := serviceTasks.CreateTask(service.Task)

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
	service := &serviceTasks.TaskService{
		Task: &domain.STask{ID: taskId},
	}
	err := service.GetSingleTask()

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Task",
		"data":    service.Task,
	})

}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	userId := c.Param("userId")

	service := &serviceTasks.TaskService{
		Task: &domain.STask{ID: taskId, UserID: userId},
	}

	err := service.DeleteTask()

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Task deleted successfully",
	})
}
