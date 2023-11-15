package models

import "time"

// Field model
type Field struct {
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	SheetID   int64     `bun:"sheet_id,notnull" json:"sheet_id"`
	Name      string    `bun:"name,type:text,notnull" json:"name"`
	Type      string    `bun:"type,notnull" json:"type"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
