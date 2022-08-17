package adapterPostgresRepo

import (
	"fmt"
	"log"
	"time"

	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

type SUserRepo struct {
	DB *pg.DB
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

// Insert new user data to database
func (userRepo SUserRepo) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := userRepo.DB.Insert(user)

	return user, err
}

// Query get all user's data in the database
func (userRepo SUserRepo) GetAllData() ([]domain.SUser, int, error) {

	users := []domain.SUser{}

	count, err := userRepo.DB.Model(&users).SelectAndCount()
	// count, err := PostgresConnect.Model(users).SelectAndCount()
	return users, count, err
}

// Query get single user's data by ID in the database
func (userRepo SUserRepo) GetSingleData(user *domain.SUser) error {
	fmt.Println("go hererrere")
	return userRepo.DB.Select(user)
}

// Query get all user's data in the database
func (userRepo SUserRepo) EditData(user *domain.SUser) error {

	_, err := userRepo.DB.Model(user).
		Set("username = ?", user.Username).
		Set(`limit_per_day = ?`, user.LimitPerDay).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", user.ID).
		Update()

	return err
}

// Query get all user's data in the database
func (userRepo SUserRepo) DeleteData(user *domain.SUser) error {

	return userRepo.DB.Delete(user)
}
