package serviceUsers

import (
	adapterPostgresRepo "togo/internal/adapter/postgressql/repositories"
	"togo/internal/core/domain"
)

// Bushiness for get all user
func GetAllUsers() ([]domain.SUser, int, error) {
	var users []domain.SUser

	count, err := adapterPostgresRepo.UserQueryGetAllData(&users)

	return users, count, err
}

// Bushiness for create new user
func CreateUser(user domain.SUser) (domain.SUser, error) {

	return adapterPostgresRepo.UserQueryCreateData(user.Username, user.Limit)
}

func GetSingleUser(user *domain.SUser) error {

	return adapterPostgresRepo.UserQueryGetSingleData(user)
}

func EditUser(userId string, user domain.SUser) error {

	return adapterPostgresRepo.UserQueryEditData(userId, user)
}

func DeleteUser(user *domain.SUser) error {

	return adapterPostgresRepo.UserQueryDeleteData(user)
}
