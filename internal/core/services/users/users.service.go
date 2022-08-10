package serviceUsers

import (
	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type UserService struct {
	DB       *pg.DB
	UserRepo interface {
		// GetAllUsers() ([]domain.SUser, int, error)
		CreateUser(db *pg.DB, user *domain.SUser) (*domain.SUser, error)
		// GetSingleUser() error
		// DeleteUser() error
		// UserQueryGetAllData() ([]*domain.SUser, int, error)
	}
}

// Bushiness for get all user
// func (service *UserService) GetAllUsers() ([]*domain.SUser, int, error) {
// 	users, count, err := service.UserRepo.UserQueryGetAllData()

// 	return users, count, err
// }

// Business for create new user
func (service *UserService) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	user.ID = uuid.New().String()

	return service.UserRepo.CreateUser(service.DB, user)
}

// func (service *UserService) GetSingleUser() error {
// 	repo := &adapterPostgresRepo.SUserRepo{
// 		User: service.User,
// 	}
// 	return repo.UserQueryGetSingleData()
// }

// func (service *UserService) EditUser() error {
// 	repo := &adapterPostgresRepo.SUserRepo{
// 		User: service.User,
// 	}
// 	return repo.UserQueryEditData()
// }

// func (service *UserService) DeleteUser() error {
// 	repo := &adapterPostgresRepo.SUserRepo{
// 		User: service.User,
// 	}
// 	return repo.UserQueryDeleteData()
// }
