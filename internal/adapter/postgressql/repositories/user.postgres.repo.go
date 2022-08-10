package adapterPostgresRepo

import (
	"log"
	"time"

	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

type SUserRepo struct {
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
func (userRepo SUserRepo) UserQueryGetAllData() ([]*domain.SUser, int, error) {

	users := []*domain.SUser{}
	count, err := PostgresConnect.Model(users).SelectAndCount()
	return users, count, err
}

// Insert new user data to database
func (userRepo SUserRepo) CreateUser(db *pg.DB, user *domain.SUser) (*domain.SUser, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := db.Insert(&user)
	return user, err
}

// Query get single user's data by ID in the database
// func (userRepo SUserRepo) UserQueryGetSingleData() error {

// 	return PostgresConnect.Select(userRepo.User)

// }

// // Query get all user's data in the database
// func (userRepo SUserRepo) UserQueryEditData() error {
// 	_, err := PostgresConnect.Model(&domain.SUser{}).Set("username = ?", userRepo.User.Username, "limit = ?", userRepo.User.Limit, "updated_at = ?", time.Now()).Where("id = ?", userRepo.User.ID).Update()

// 	return err
// }

// // Query get all user's data in the database
// func (userRepo SUserRepo) UserQueryDeleteData() error {

// 	return PostgresConnect.Delete(userRepo.User)
// }
