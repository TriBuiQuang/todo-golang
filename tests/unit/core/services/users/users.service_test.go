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

type mockUserRepo struct {
	err error
}

func (m *mockUserRepo) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	return user, nil
}
func (m *mockUserRepo) GetAllData() ([]domain.SUser, int, error) {
	return nil, -1, nil
}
func (m *mockUserRepo) GetSingleData(user *domain.SUser) error {
	return nil
}
func (m *mockUserRepo) EditData(user *domain.SUser) error {
	return nil
}
func (m *mockUserRepo) DeleteData(user *domain.SUser) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo := new(MockRepository)
	mockRepo.On("CreateUser").Return(user, nil)

	userService := &serviceUsers.UserService{UserRepo: &mockUserRepo{}}

	// Run function
	result, err := userService.CreateUser(&user)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.LimitPerDay, result.LimitPerDay)
	assert.Equal(t, nil, err)

}
