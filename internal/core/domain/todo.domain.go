package domain

import "time"

// STodo represents data about a record todo.
type STodo struct {
	tableName struct{}  `pg:"todos"`         // default values are the same
	ID        string    `pg:",pk" json:"id"` // ID is primary key
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt time.Time `pg:"default:now()" json:"updated_at"`
}
