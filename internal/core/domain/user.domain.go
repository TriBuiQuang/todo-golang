package domain

import "time"

// album represents data about a record album.
type SUser struct {
	tableName struct{}  `pg:"users"`                   // default values are the same
	ID        string    `pg:",pk" json:"id"`           // ID is primary key
	Username  string    `pg:",unique" json:"username"` // username is unique
	Limit     int       `pg:"default:10" json:"limit"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt time.Time `pg:"default:now()" json:"updated_at"`
}
