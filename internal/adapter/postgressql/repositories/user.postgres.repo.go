package adapterPostgresRepo

import (
	"log"
	"time"

	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

// Giao tiep voi db (nhu may cau insert)

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
func UserQueryGetAllData(users *[]domain.SUser) (int, error) {

	return PostgresConnect.Model(users).SelectAndCount()
}

// Insert new user data to database
func UserQueryCreateData(username string, limit int) (domain.SUser, error) {
	id := guuid.New().String()
	now := time.Now()

	newUser := domain.SUser{
		ID:        id,
		Username:  username,
		Limit:     limit,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := PostgresConnect.Insert(&newUser)

	return newUser, err
}

// Query get single user's data by ID in the database
func UserQueryGetSingleData(user *domain.SUser) error {

	return PostgresConnect.Select(user)

}

// Query get all user's data in the database
func UserQueryEditData(userId string, user domain.SUser) error {
	_, err := PostgresConnect.Model(&domain.SUser{}).Set("username = ?", user.Username, "limit = ?", user.Limit, "updated_at = ?", time.Now()).Where("id = ?", userId).Update()

	return err
}

// Query get all user's data in the database
func UserQueryDeleteData(user *domain.SUser) error {

	return PostgresConnect.Delete(user)
}
