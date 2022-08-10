package domain

import "time"

// album represents data about a record album.
type STask struct {
	tableName struct{}  `pg:"tasks"`         // default values are the same
	ID        string    `pg:",pk" json:"id"` // ID is primary key
	UserID    string    `pg:"rel:has-one" json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt time.Time `pg:"default:now()" json:"updated_at"`
}

func (STask) TableName() string {
	return "tasks"
}
