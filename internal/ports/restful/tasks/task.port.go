package portsRestFulTask

import (
	"net/http"
	adapterPostgresRepo "togo/internal/adapter/postgresql/repositories"
	"togo/internal/core/domain"
	serviceTasks "togo/internal/core/services/tasks"
	serviceUsers "togo/internal/core/services/users"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

type STaskPort struct {
	TaskService interface {
		CreateTask(task *domain.STask) (*domain.STask, error)
		GetAllTasks() ([]domain.STask, int, error)
		GetAllTaskToday(tasks *[]domain.STask, userId string) (*[]domain.STask, int, error)
		GetSingleTask(task *domain.STask) error
		DeleteTask(task *domain.STask) error
	}
}

func NewTaskPort(db *pg.DB) *STaskPort {
	taskService := &serviceTasks.TaskService{
		DB: db,
		TaskRepo: &adapterPostgresRepo.STaskRepo{
			DB: db,
		},
		UserService: serviceUsers.UserService{
			DB: db,
			UserRepo: &adapterPostgresRepo.SUserRepo{
				DB: db,
			},
		},
	}

	return &STaskPort{
		TaskService: taskService,
	}
}

func (t *STaskPort) CreateTask(c *gin.Context) {
	task := &domain.STask{}
	c.BindJSON(task)

	if portsRestFul.ValidateCreateTask(c, task) {
		return
	}

	newTask, err := t.TaskService.CreateTask(task)

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

func (t *STaskPort) GetAllTasks(c *gin.Context) {
	tasks, totalTasks, err := t.TaskService.GetAllTasks()

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

func (t *STaskPort) GetSingleTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task := &domain.STask{ID: taskId}

	err := t.TaskService.GetSingleTask(task)

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

func (t *STaskPort) DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	userId := c.Param("userId")
	task := &domain.STask{ID: taskId, UserID: userId}

	err := t.TaskService.DeleteTask(task)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Task deleted successfully",
	})
}
