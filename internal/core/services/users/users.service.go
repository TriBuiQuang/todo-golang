package serviceUsers

import (
	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type UserService struct {
	DB       *pg.DB
	UserRepo interface {
		CreateUser(user *domain.SUser) (*domain.SUser, error)
		GetAllData() ([]*domain.SUser, int, error)
		GetSingleData(user *domain.SUser) error
		EditData(user *domain.SUser) error
		DeleteData(user *domain.SUser) error
	}
}

// Business for create new user
func (service *UserService) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	user.ID = uuid.New().String()

	return service.UserRepo.CreateUser(user)
}

// Bushiness for get all user
func (service *UserService) GetAllUsers() ([]*domain.SUser, int, error) {

	return service.UserRepo.GetAllData()
}

func (service *UserService) GetSingleUser(user *domain.SUser) error {

	return service.UserRepo.GetSingleData(user)
}

func (service *UserService) EditUser(user *domain.SUser) error {

	return service.UserRepo.EditData(user)
}

func (service *UserService) DeleteUser(user *domain.SUser) error {

	return service.UserRepo.DeleteData(user)
}
