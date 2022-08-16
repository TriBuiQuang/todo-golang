package portsRestFulTask

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	mockRestFul "togo/__mocks__/ports/restful"
	"togo/internal/core/domain"
	portsRestFulTask "togo/internal/ports/restful/tasks"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/assert"
)

type mockTaskService struct {
	err error
}

func (m *mockTaskService) CreateTask(task *domain.STask) (*domain.STask, error) {
	return nil, m.err
}
func (m *mockTaskService) GetAllTasks() ([]domain.STask, int, error) {
	return nil, -1, m.err
}
func (m *mockTaskService) GetAllTaskToday(tasks *[]domain.STask, userId string) (*[]domain.STask, int, error) {
	return nil, -1, m.err
}
func (m *mockTaskService) GetSingleTask(task *domain.STask) error {
	return m.err
}
func (m *mockTaskService) DeleteTask(task *domain.STask) error {
	return m.err
}

func TestNewTaskPort(t *testing.T) {

	result := portsRestFulTask.NewTaskPort(&pg.DB{})

	assert.Equal(t, "*portsRestFulTask.STaskPort", reflect.TypeOf(result).String())
	assert.Equal(t, "*serviceTasks.TaskService", reflect.TypeOf(result.TaskService).String())
}

func TestCreateTask(t *testing.T) {
	t.Parallel()
	taskService := &mockTaskService{}
	testCases := []struct {
		name        string
		req         *domain.STask
		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",
			req:  &domain.STask{Name: "testing case", UserID: "123213"},
			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusCreated,
			actualErr:  nil,
		},
		{
			name: "fail because missing Name's task",
			req:  &domain.STask{UserID: "123213"},
			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because missing UserId",
			req:  &domain.STask{Name: "testing case"},
			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because missing Name and UserId",
			req:  &domain.STask{},
			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because give normal json but fail in service layer about over the limit today",
			req:  &domain.STask{Name: "testing case", UserID: "123213"},
			setup: func() {
				taskService = &mockTaskService{
					err: errors.New("this user is out of the limit in order to create a new task. "),
				}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because give normal json but fail in service layer",
			req:  &domain.STask{Name: "testing case", UserID: "123213"},
			setup: func() {
				taskService = &mockTaskService{
					err: errors.New("test fail"),
				}
			},
			actualCode: http.StatusInternalServerError,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			taskPort := &portsRestFulTask.STaskPort{
				TaskService: taskService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			mockRestFul.MockJsonPost(c, testCase.req)
			taskPort.CreateTask(c)

			assert.Equal(t, response.Code, testCase.actualCode)
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	t.Parallel()
	taskService := &mockTaskService{}
	testCases := []struct {
		name        string
		req         *domain.STask
		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",
			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusOK,
			actualErr:  nil,
		},
		{
			name: "fail because give normal json but fail in service layer",
			setup: func() {
				taskService = &mockTaskService{
					err: errors.New("test fail"),
				}
			},
			actualCode: http.StatusInternalServerError,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			taskPort := &portsRestFulTask.STaskPort{
				TaskService: taskService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			taskPort.GetAllTasks(c)

			assert.Equal(t, response.Code, testCase.actualCode)
		})
	}
}

func TestGetSingleTask(t *testing.T) {
	t.Parallel()
	taskService := &mockTaskService{}
	testCases := []struct {
		name        string
		params      []gin.Param
		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",

			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusOK,
			actualErr:  nil,
		},
		{
			name: "fail because give normal params but fail in service layer",

			setup: func() {
				taskService = &mockTaskService{
					err: errors.New("test fail"),
				}
			},
			actualCode: http.StatusNotFound,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			taskPort := &portsRestFulTask.STaskPort{
				TaskService: taskService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			c.Params = testCase.params
			taskPort.GetSingleTask(c)

			assert.Equal(t, testCase.actualCode, response.Code)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	t.Parallel()
	taskService := &mockTaskService{}
	testCases := []struct {
		name        string
		params      []gin.Param
		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",

			setup: func() {
				taskService = &mockTaskService{}
			},
			actualCode: http.StatusOK,
			actualErr:  nil,
		},
		{
			name: "fail because give normal params but fail in service layer",

			setup: func() {
				taskService = &mockTaskService{
					err: errors.New("test fail"),
				}
			},
			actualCode: http.StatusInternalServerError,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			taskPort := &portsRestFulTask.STaskPort{
				TaskService: taskService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			c.Params = testCase.params
			taskPort.DeleteTask(c)

			assert.Equal(t, testCase.actualCode, response.Code)
		})
	}
}
