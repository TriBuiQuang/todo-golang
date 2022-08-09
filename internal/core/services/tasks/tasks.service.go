package serviceTasks

// func CreateNewTask

import (
	"errors"
	"time"
	adapterPostgresRepo "togo/internal/adapter/postgressql/repositories"
	"togo/internal/core/domain"
)

type TaskService struct {
	TaskRepo interface {
		GetAllTasks() ([]domain.STask, int, error)
	}
}

func (service *TaskService) GetAllTasks() ([]domain.STask, int, error) {
	var tasks []domain.STask

	repo := &adapterPostgresRepo.STaskRepo{
		Tasks: &tasks,
	}

	count, err := repo.TaskQueryGetAllData()

	return tasks, count, err
}

// Bushiness for get all task
func GetAllTasks() ([]domain.STask, int, error) {
	var tasks []domain.STask

	count, err := adapterPostgresRepo.TaskQueryGetAllData(&tasks)

	return tasks, count, err
}

// Bushiness for create new task
func CreateTask(task *domain.STask) (domain.STask, error) {
	var oldTask []domain.STask
	currentUser := &domain.SUser{ID: task.UserID}
	now := time.Now()

	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	beginningOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	userErr := adapterPostgresRepo.UserQueryGetSingleData(currentUser)
	if userErr != nil {
		return domain.STask{}, userErr
	}

	totalTodayTask, oldTaskErr := adapterPostgresRepo.GetAllTaskToday(&oldTask, task.UserID, beginningOfDay)

	if oldTaskErr != nil {
		return domain.STask{}, oldTaskErr
	}

	if totalTodayTask != 0 && totalTodayTask >= currentUser.Limit {
		return *task, errors.New("this user is out of the limit in order to create a new task. ")
	}

	return adapterPostgresRepo.TaskQueryCreateData(task)
}

func GetSingleTask(task *domain.STask) error {

	return adapterPostgresRepo.TaskQueryGetSingleData(task)
}

func DeleteTask(task *domain.STask) error {

	return adapterPostgresRepo.TaskQueryDeleteData(task)
}
