package adapterPostgresRepo

import (
	"log"
	"time"

	"togo/internal/core/domain"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

type STaskRepo struct {
	DB *pg.DB
}

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

// Insert new task data to database
func (taskRepo STaskRepo) CreateData(task *domain.STask) (*domain.STask, error) {
	// id := guuid.New().String()
	// now := time.Now()

	// newData := domain.STask{
	// 	ID:        id,
	// 	UserID:    taskRepo.Task.UserID,
	// 	Name:      taskRepo.Task.Name,
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }

	// err := PostgresConnect.Insert(&newData)

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	err := taskRepo.DB.Insert(task)

	return task, err
}

// Query get all task's data in the database
func (taskRepo STaskRepo) GetAllData() ([]domain.STask, int, error) {
	tasks := []domain.STask{}

	count, err := taskRepo.DB.Model(&tasks).SelectAndCount()

	return tasks, count, err
}

func (taskRepo STaskRepo) GetAllTaskToday(tasks *[]domain.STask, userId string, beginningOfDay time.Time) (int, error) {

	return taskRepo.DB.Model(tasks).Where("user_id = ?", userId).Where("created_at >= ?", beginningOfDay).SelectAndCount()
}

// Query get single task's data by ID in the database
func (taskRepo STaskRepo) GetSingleData(task *domain.STask) error {

	return taskRepo.DB.Select(task)
}

// Query get all task's data in the database
func (taskRepo STaskRepo) DeleteData(task *domain.STask) error {

	return PostgresConnect.Delete(task)
}
