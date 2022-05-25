package models

import (
	"time"
)

type Product struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,lte=255"`
	Active    bool      `gorm:"default:true" db:"is_active" json:"is_active"`
	IconID    uint32    `gorm:"default:null" db:"icon_id" json:"icon_id"`
	Icon      File      `json:"icon"`
	ImageID   uint32    `db:"image_id" json:"image_id"`
	Image     File      `json:"image"`
	PageID    uint32    `gorm:"default:null" db:"page_id" json:"page_id"`
	Page      Page      `json:"page"`
}
