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

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
// var dbConnect *pg.DB

// func InitiateDB(db *pg.DB) {
// 	dbConnect = db
// }

// Create User Table
func TodoCreateTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{IfNotExists: true}

	createError := db.CreateTable(&domain.STodo{}, opts)

	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Todo table created")

	return nil
}

// Query get all todo's data in the database
func TodoQueryGetAllData(todos *[]domain.STodo) error {
	err := PostgresConnect.Model(todos).Select()

	return err
}

// Insert new todo data to database
func TodoQueryCreateData(title, body, completed string) (domain.STodo, error) {
	id := guuid.New().String()
	now := time.Now()

	newData := domain.STodo{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := PostgresConnect.Insert(&newData)

	return newData, err
}

// Query get single todo's data by ID in the database
func TodoQueryGetSingleData(todo *domain.STodo) error {
	err := PostgresConnect.Select(todo)

	return err
}

// Query get all todo's data in the database
func TodoQueryEditData(todoId, completed string) error {
	_, err := PostgresConnect.Model(&domain.STodo{}).Set("completed = ?", completed, "updated_at = ?", time.Now()).Where("id = ?", todoId).Update()

	return err
}

// Query get all todo's data in the database
func TodoQueryDeleteData(todo *domain.STodo) error {

	return PostgresConnect.Delete(todo)
}
