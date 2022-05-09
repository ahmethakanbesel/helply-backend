package models

import (
	"time"
)

type Product struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,lte=255"`
	Icon      string    `gorm:"type:varchar(255)" db:"icon" json:"icon" validate:"lte=255"`
	Image     string    `gorm:"type:varchar(255)" db:"image" json:"image" validate:"lte=255"`
	PageID    uint32    `db:"page_id" json:"page_id"`
	Page      Page
}
