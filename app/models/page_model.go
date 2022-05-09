package models

import (
	"time"
)

type Page struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Title     string    `gorm:"type:varchar(255)" db:"title" json:"title" validate:"required,lte=255"`
	Slug      string    `gorm:"type:varchar(255)" db:"slug" json:"slug" validate:"lte=255"`
	Content   string    `db:"content" json:"content"`
}
