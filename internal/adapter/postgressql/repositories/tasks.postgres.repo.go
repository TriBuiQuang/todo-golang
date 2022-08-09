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
func TaskCreateTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{IfNotExists: true}

	createError := db.CreateTable(&domain.STask{}, opts)

	if createError != nil {
		log.Printf("Error while creating task table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Task table created")

	return nil
}

// Query get all task's data in the database
func TaskQueryGetAllData(tasks *[]domain.STask) (int, error) {

	return PostgresConnect.Model(tasks).SelectAndCount()
}

func GetAllTaskToday(tasks *[]domain.STask, userId string, beginningOfDay time.Time) (int, error) {

	return PostgresConnect.Model(tasks).Where("user_id = ?", userId).Where("created_at >= ?", beginningOfDay).SelectAndCount()
}

// Insert new task data to database
func TaskQueryCreateData(task *domain.STask) (domain.STask, error) {
	id := guuid.New().String()
	now := time.Now()

	newData := domain.STask{
		ID:        id,
		UserID:    task.UserID,
		Name:      task.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := PostgresConnect.Insert(&newData)

	return newData, err
}

// Query get single task's data by ID in the database
func TaskQueryGetSingleData(task *domain.STask) error {
	err := PostgresConnect.Select(task)

	return err
}

// Query get all task's data in the database
func TaskQueryDeleteData(task *domain.STask) error {

	return PostgresConnect.Delete(task)
}
