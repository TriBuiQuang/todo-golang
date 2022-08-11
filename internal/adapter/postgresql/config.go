package adapterPostgres

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"

	// serviceTasks "togo/internal/core/services/tasks"
	adapterPostgresRepo "togo/internal/adapter/postgresql/repositories"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "username"
	password = "123"
	dbname   = "test"
)

// Connecting to db
func ConnectDatabase() *pg.DB {
	opts := &pg.Options{
		User:     user,
		Password: password,
		Addr:     host + ":" + port,
		Database: dbname,
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect postgres database")
		os.Exit(100)
	}

	log.Printf("Connected to db")
	adapterPostgresRepo.CreateAllTable(db)
	adapterPostgresRepo.InitiateDB(db)
	return db
}
