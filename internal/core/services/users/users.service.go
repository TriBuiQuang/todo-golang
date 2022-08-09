package serviceUsers

import (
	adapterPostgresRepo "togo/internal/adapter/postgressql/repositories"
	"togo/internal/core/domain"
)

type UserService struct {
	UserRepo interface {
		GetAllUsers() ([]domain.SUser, int, error)
		CreateUser() (domain.SUser, error)
		GetSingleUser() error
		DeleteUser() error
	}
	// Handle more than 1 data
	Users *[]domain.SUser
	// Handle only 1 data
	User *domain.SUser
}

// Bushiness for get all user
func (service *UserService) GetAllUsers() ([]domain.SUser, int, error) {
	repo := &adapterPostgresRepo.SUserRepo{
		Users: &[]domain.SUser{},
	}
	count, err := repo.UserQueryGetAllData()

	return *repo.Users, count, err
}

// Bushiness for create new user
func (service *UserService) CreateUser() (domain.SUser, error) {
	repo := &adapterPostgresRepo.SUserRepo{
		User: service.User,
	}

	return repo.UserQueryCreateData()
}

func (service *UserService) GetSingleUser() error {
	repo := &adapterPostgresRepo.SUserRepo{
		User: service.User,
	}
	return repo.UserQueryGetSingleData()
}

func (service *UserService) EditUser() error {
	repo := &adapterPostgresRepo.SUserRepo{
		User: service.User,
	}
	return repo.UserQueryEditData()
}

func (service *UserService) DeleteUser() error {
	repo := &adapterPostgresRepo.SUserRepo{
		User: service.User,
	}
	return repo.UserQueryDeleteData()
}
