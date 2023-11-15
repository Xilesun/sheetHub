package models

import "time"

// Sheet model
type Sheet struct {
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Name      string    `bun:"name,type:text,notnull" json:"name"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
