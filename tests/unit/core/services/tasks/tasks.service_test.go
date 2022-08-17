package serviceTasks

import (
	"errors"
	"testing"
	"time"
	mockServiceUsers "togo/__mocks__/core/services"
	"togo/internal/core/domain"
	serviceTasks "togo/internal/core/services/tasks"
	serviceUsers "togo/internal/core/services/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type mockTaskRepo struct {
	err error
}

func (m *mockTaskRepo) CreateData(task *domain.STask) (*domain.STask, error) {
	return task, m.err
}
func (m *mockTaskRepo) GetAllData() ([]domain.STask, int, error) {
	return nil, -1, m.err
}
func (m *mockTaskRepo) GetAllTaskToday(task *[]domain.STask, userId string, beginningOfDay time.Time) (int, error) {

	return len(*task), m.err
}
func (m *mockTaskRepo) GetSingleData(task *domain.STask) error {
	return m.err
}
func (m *mockTaskRepo) DeleteData(task *domain.STask) error {
	return m.err
}

func TestCreateTask(t *testing.T) {
	t.Parallel()

	now := time.Now()
	mockRepo := new(MockRepository)

	taskRepo := &mockTaskRepo{}
	userRepo := &mockServiceUsers.MockUserRepo{}
	testCases := []struct {
		name           string
		user           *domain.SUser
		task           *domain.STask
		setup          func()
		expectTaskName string
		expectNotID    string
		expectErr      error
	}{
		{
			name: "success normal case",
			user: &domain.SUser{
				ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now,
			},
			task: &domain.STask{ID: "123", Name: "testing task", UserID: "123"},
			setup: func() {
				taskRepo = &mockTaskRepo{}
				userRepo = &mockServiceUsers.MockUserRepo{}
			},
			expectTaskName: "testing task",
			expectNotID:    "123",
			expectErr:      nil,
		},
		{
			name: "fail because give normal json but fail in repo user layer",
			user: &domain.SUser{
				ID: "123", Username: "testing username", LimitPerDay: 1, UpdatedAt: now, CreatedAt: now,
			},
			task: &domain.STask{ID: "123", Name: "testing task", UserID: "123"},
			setup: func() {
				taskRepo = &mockTaskRepo{}
				userRepo = &mockServiceUsers.MockUserRepo{
					Err: errors.New("fail test in user repo"),
				}
			},
			expectTaskName: "",
			expectNotID:    "123",
			expectErr:      errors.New("fail test in user repo"),
		},
		{
			name: "fail because give normal json but fail in repo task layer",
			user: &domain.SUser{
				ID: "123", Username: "testing username", LimitPerDay: 0, UpdatedAt: now, CreatedAt: now,
			},
			task: &domain.STask{ID: "123", Name: "testing task", UserID: "123"},
			setup: func() {
				taskRepo = &mockTaskRepo{
					err: errors.New("fail test in task repo"),
				}
				userRepo = &mockServiceUsers.MockUserRepo{}
			},
			expectTaskName: "",
			expectNotID:    "123",
			expectErr:      errors.New("fail test in task repo"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()

			mockRepo.On("CreateTask").Return(testCase.task, nil)
			mockRepo.On("GetAllTaskToday").Return(testCase.task, nil)
			mockRepo.On("GetSingleData").Return(testCase.user, nil)

			taskService := &serviceTasks.TaskService{
				TaskRepo:    taskRepo,
				UserService: serviceUsers.UserService{UserRepo: userRepo},
			}

			// Run function
			result, err := taskService.CreateTask(testCase.task)

			assert.Equal(t, testCase.expectErr, err)
			assert.Equal(t, testCase.expectTaskName, result.Name)
			assert.NotEqual(t, testCase.expectNotID, result.ID)
		})
	}
}
