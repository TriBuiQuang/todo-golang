package adapterPostgresRepo

import (
	"log"
	"time"

	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

type SUserRepo struct {
	TaskInterface interface {
		UserQueryGetAllData() (int, error)
		UserQueryCreateData() (domain.SUser, error)
		UserQueryGetSingleData() error
		UserQueryEditData() error
		UserQueryDeleteData() error
	}
	// Handle more than 1 data
	Users *[]domain.SUser
	// Handle only 1 data
	User *domain.SUser
}

// Create User Table
func UserCreateTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{IfNotExists: true}

	createError := db.CreateTable(&domain.SUser{}, opts)

	if createError != nil {
		log.Printf("Error while creating user table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("User table created")

	return nil
}

// Query get all user's data in the database
func (userRepo SUserRepo) UserQueryGetAllData() (int, error) {

	return PostgresConnect.Model(userRepo.Users).SelectAndCount()
}

// Insert new user data to database
func (userRepo SUserRepo) UserQueryCreateData() (domain.SUser, error) {
	id := guuid.New().String()
	now := time.Now()

	newUser := domain.SUser{
		ID:        id,
		Username:  userRepo.User.Username,
		Limit:     userRepo.User.Limit,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := PostgresConnect.Insert(&newUser)

	return newUser, err
}

// Query get single user's data by ID in the database
func (userRepo SUserRepo) UserQueryGetSingleData() error {

	return PostgresConnect.Select(userRepo.User)

}

// Query get all user's data in the database
func (userRepo SUserRepo) UserQueryEditData() error {
	_, err := PostgresConnect.Model(&domain.SUser{}).Set("username = ?", userRepo.User.Username, "limit = ?", userRepo.User.Limit, "updated_at = ?", time.Now()).Where("id = ?", userRepo.User.ID).Update()

	return err
}

// Query get all user's data in the database
func (userRepo SUserRepo) UserQueryDeleteData() error {

	return PostgresConnect.Delete(userRepo.User)
}
