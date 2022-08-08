package adapterPostgresRepo

import (
	"github.com/go-pg/pg/v9"
)

// Giao tiep voi db (nhu may cau insert)

var PostgresConnect *pg.DB

func InitiateDB(db *pg.DB) {
	PostgresConnect = db
}

func CreateAllTable(db *pg.DB) error {
	UserCreateTable(db)
	TaskCreateTable(db)
	TodoCreateTable(db)
	return nil
}
