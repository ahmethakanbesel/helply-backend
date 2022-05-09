package models

import (
	"time"
)

type Language struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Code      string    `gorm:"type:varchar(7)" db:"code" json:"code" validate:"required,lte=4"`
	Name      string    `gorm:"type:varchar(31)" db:"name" json:"name" validate:"required,lte=255"`
}

type LanguageString struct {
	ID          uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Slug        string    `gorm:"type:varchar(127);uniqueIndex" db:"slug" json:"slug" validate:"required,lte=127"`
	LanguageID  uint32    `db:"language_id" json:"language_id" validate:"required"`
	Language    Language
	Translation string `db:"translation" json:"translation" validate:"required"`
}
