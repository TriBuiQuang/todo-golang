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
		CreateTask() (domain.STask, error)
		GetSingleTask() error
		DeleteTask() error
	}
	// Handle more than 1 data
	Tasks *[]domain.STask
	// Handle only 1 data
	Task *domain.STask
}

func (service *TaskService) GetAllTasks() ([]domain.STask, int, error) {

	repo := &adapterPostgresRepo.STaskRepo{
		Tasks: &[]domain.STask{},
	}

	count, err := repo.TaskQueryGetAllData()

	return *repo.Tasks, count, err
}

// Bushiness for create new task
func (service *TaskService) CreateTask() (domain.STask, error) {
	var oldTask []domain.STask

	repo := &adapterPostgresRepo.STaskRepo{
		Task: &domain.STask{},
	}
	repoUser := &adapterPostgresRepo.SUserRepo{
		User: &domain.SUser{ID: service.Task.UserID},
	}

	now := time.Now()

	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	beginningOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	userErr := repoUser.UserQueryGetSingleData()
	if userErr != nil {
		return domain.STask{}, userErr
	}

	totalTodayTask, oldTaskErr := adapterPostgresRepo.GetAllTaskToday(&oldTask, repo.Task.UserID, beginningOfDay)

	if oldTaskErr != nil {
		return domain.STask{}, oldTaskErr
	}

	if totalTodayTask != 0 && totalTodayTask >= repoUser.User.Limit {
		return *repo.Task, errors.New("this user is out of the limit in order to create a new task. ")
	}

	return adapterPostgresRepo.TaskQueryCreateData(repo.Task)
}

func (service *TaskService) GetSingleTask() error {

	return adapterPostgresRepo.TaskQueryGetSingleData(service.Task)
}

func (service *TaskService) DeleteTask() error {

	return adapterPostgresRepo.TaskQueryDeleteData(service.Task)
}
