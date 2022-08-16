package serviceUsers

import (
	"testing"
	"time"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type MockUserRepo struct {
	err error
}

func (m *MockUserRepo) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	return user, nil
}
func (m *MockUserRepo) GetAllData() ([]domain.SUser, int, error) {
	return []domain.SUser{}, -1, nil
}
func (m *MockUserRepo) GetSingleData(user *domain.SUser) error {
	return nil
}
func (m *MockUserRepo) EditData(user *domain.SUser) error {
	return nil
}
func (m *MockUserRepo) DeleteData(user *domain.SUser) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo := new(MockRepository)
	mockRepo.On("CreateUser").Return(user, nil)

	userService := &serviceUsers.UserService{UserRepo: &MockUserRepo{}}

	// Run function
	result, err := userService.CreateUser(&user)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.LimitPerDay, result.LimitPerDay)
	assert.Equal(t, nil, err)
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo := new(MockRepository)
	mockRepo.On("GetAllData").Return(user, nil)

	userService := &serviceUsers.UserService{UserRepo: &MockUserRepo{}}

	// Run function
	_, total, err := userService.GetAllUsers()

	assert.Equal(t, -1, total)
	assert.Equal(t, nil, err)
}

func TestGetSingleUser(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo := new(MockRepository)
	mockRepo.On("GetSingleUser").Return(user, nil)

	userService := &serviceUsers.UserService{UserRepo: &MockUserRepo{}}

	// Run function
	err := userService.GetSingleUser(&user)
	assert.Equal(t, nil, err)
}

func TestEditUser(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo := new(MockRepository)
	mockRepo.On("EditUser").Return(user, nil)

	userService := &serviceUsers.UserService{UserRepo: &MockUserRepo{}}

	// Run function
	err := userService.EditUser(&user)
	assert.Equal(t, nil, err)
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo := new(MockRepository)
	mockRepo.On("DeleteUser").Return(user, nil)

	userService := &serviceUsers.UserService{UserRepo: &MockUserRepo{}}

	// Run function
	err := userService.DeleteUser(&user)
	assert.Equal(t, nil, err)
}
