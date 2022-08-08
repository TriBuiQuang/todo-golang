package serviceTasks

// func CreateNewTask

import (
	adapterPostgresRepo "togo/internal/adapter/postgressql/repositories"
	"togo/internal/core/domain"
)

// Bushiness for get all task
func GetAllTasks() ([]domain.STask, int, error) {
	var tasks []domain.STask

	count, err := adapterPostgresRepo.TaskQueryGetAllData(&tasks)

	return tasks, count, err
}

// Bushiness for create new task
func CreateTask(task domain.STask) (domain.STask, error) {

	return adapterPostgresRepo.TaskQueryCreateData(task)
}

func GetSingleTask(task *domain.STask) error {

	return adapterPostgresRepo.TaskQueryGetSingleData(task)
}

func EditTask(taskId string, task domain.STask) error {

	return adapterPostgresRepo.TaskQueryEditData(taskId, task)
}

func DeleteTask(task *domain.STask) error {

	return adapterPostgresRepo.TaskQueryDeleteData(task)
}
