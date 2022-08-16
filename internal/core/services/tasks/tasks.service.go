package serviceTasks

// func CreateNewTask

import (
	"errors"
	"fmt"
	"time"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type TaskService struct {
	DB       *pg.DB
	TaskRepo interface {
		CreateData(task *domain.STask) (*domain.STask, error)
		GetAllData() ([]domain.STask, int, error)
		GetAllTaskToday(task *[]domain.STask, userId string, beginningOfDay time.Time) (int, error)
		GetSingleData(task *domain.STask) error
		DeleteData(task *domain.STask) error
	}
	UserService serviceUsers.UserService
}

// Bushiness for create new task
func (service *TaskService) CreateTask(task *domain.STask) (*domain.STask, error) {
	var oldTask []domain.STask
	task.ID = uuid.New().String()

	user := &domain.SUser{ID: task.UserID}
	now := time.Now()

	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	beginningOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	userErr := service.UserService.UserRepo.GetSingleData(user)

	if userErr != nil {
		return &domain.STask{}, userErr
	}
	fmt.Println("asdasdasd : ", user)

	totalTodayTask, oldTaskErr := service.TaskRepo.GetAllTaskToday(&oldTask, task.UserID, beginningOfDay)

	if oldTaskErr != nil {
		return &domain.STask{}, oldTaskErr
	}

	if totalTodayTask != 0 && totalTodayTask >= user.LimitPerDay {
		return &domain.STask{}, errors.New("this user is out of the limit in order to create a new task. ")
	}

	return service.TaskRepo.CreateData(task)
}

func (service *TaskService) GetAllTasks() ([]domain.STask, int, error) {

	return service.TaskRepo.GetAllData()
}

func (service *TaskService) GetAllTaskToday(tasks *[]domain.STask, userId string) (*[]domain.STask, int, error) {
	now := time.Now()

	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	beginningOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	count, err := service.TaskRepo.GetAllTaskToday(tasks, userId, beginningOfDay)

	return tasks, count, err
}

func (service *TaskService) GetSingleTask(task *domain.STask) error {

	return service.TaskRepo.GetSingleData(task)
}

func (service *TaskService) DeleteTask(task *domain.STask) error {

	return service.TaskRepo.DeleteData(task)
}
