package mockServiceUsers

import (
	"togo/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

type MockUserRepo struct {
	Err error
}

func (m *MockUserRepo) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	return user, m.Err
}
func (m *MockUserRepo) GetAllData() ([]*domain.SUser, int, error) {
	return []*domain.SUser{}, -1, m.Err
}
func (m *MockUserRepo) GetSingleData(user *domain.SUser) error {
	return m.Err
}
func (m *MockUserRepo) EditData(user *domain.SUser) error {
	return m.Err
}
func (m *MockUserRepo) DeleteData(user *domain.SUser) error {
	return m.Err
}
