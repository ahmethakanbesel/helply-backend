package models

import (
	"time"
)

// User struct to describe User object.
type User struct {
	ID           uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `gorm:"type:varchar(255)" db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `gorm:"type:varchar(256)" db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=256"`
	UserStatus   int       `gorm:"default:1" db:"user_status" json:"user_status" validate:"required,len=1"`
	UserRoleID   uint32    `db:"role_id" json:"role_id" validate:"required"`
	UserRole     UserRole
}

type UserRole struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `gorm:"type:varchar(31)" db:"name" json:"name" validate:"required,lte=31"`
}
