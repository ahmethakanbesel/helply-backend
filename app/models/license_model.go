package models

import (
	"time"
)

type License struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	ExpiresAt time.Time `gorm:"timestamp with time zone" db:"expires_at" json:"expires_at"`
	Code      string    `gorm:"type:varchar(255);uniqueIndex" db:"code" json:"code" validate:"required,lte=255"`
	ProductID uint32    `db:"product_id" json:"product_id" validate:"required"`
	Product   Product   `json:"product"`
	IsActive  bool      `gorm:"default:true" db:"is_active" json:"is_active"`
}

type CustomerLicense struct {
	ID         uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	CustomerID uint32    `db:"customer_id" json:"customer_id" validate:"required"`
	Customer   User
	LicenseID  uint32 `db:"license_id" json:"license_id" validate:"required"`
	License    License
}
