package serviceTodo

import (
	adapterPostgresRepo "togo/internal/adapter/postgresql/repositories"
	"togo/internal/core/domain"
)

// Bushiness for get all todo
func GetAllTodos() ([]domain.STodo, error) {
	var todos []domain.STodo

	err := adapterPostgresRepo.TodoQueryGetAllData(&todos)

	return todos, err
}

// Bushiness for create new todo
func CreateTodo(todo domain.STodo) (domain.STodo, error) {

	return adapterPostgresRepo.TodoQueryCreateData(todo.Title, todo.Body, todo.Completed)
}

func GetSingleTodo(todo *domain.STodo) error {

	return adapterPostgresRepo.TodoQueryGetSingleData(todo)
}

func EditTodo(todoId, completed string) error {

	return adapterPostgresRepo.TodoQueryEditData(todoId, completed)
}

func DeleteTodo(todo *domain.STodo) error {

	return adapterPostgresRepo.TodoQueryDeleteData(todo)
}
