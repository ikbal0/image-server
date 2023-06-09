package entities

import "time"

type GormModel struct {
	ID        uint       `gorm:"PrimaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
