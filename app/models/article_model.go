package models

import (
	"time"
)

type Article struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Title     string    `gorm:"type:varchar(255)" db:"title" json:"title" validate:"required,lte=255"`
	//Slug       string    `gorm:"type:varchar(255)" db:"slug" json:"slug" validate:"required,lte=255"`
	Content    string `db:"content" json:"content"`
	Votes      int32  `gorm:"default:0" db:"votes" json:"votes"`
	ProductID  uint32 `db:"product_id" json:"product_id" validate:"required"`
	Product    Product
	CategoryID uint32 `db:"category_id" json:"category" validate:"required"`
	Category   ArticleCategory
}

type ArticleCategory struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,lte=255"`
}

type ArticleTag struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	ArticleID uint32    `db:"article_id" json:"article_id" validate:"required"`
	Article   Article
	Name      string `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,lte=255"`
}
