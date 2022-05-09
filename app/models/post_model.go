package models

import (
	"time"
)

type Post struct {
	ID         uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UserID     uint32    `db:"user_id" json:"user_id" validate:"required"`
	User       User
	Title      string `gorm:"type:varchar(255);default:Test" db:"title" json:"title" validate:"required,lte=255"`
	Author     string `db:"author" json:"author" validate:"required,lte=255"`
	BookStatus uint8  `db:"book_status" json:"book_status" validate:"required,len=1" autoIncrement:"true"`
	Content    string `gorm:"type:text" json:"content"`
}
