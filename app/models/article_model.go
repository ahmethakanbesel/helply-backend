package models

import (
	"time"
)

type Article struct {
	ID         uint32          `db:"id" json:"id" validate:"required"`
	CreatedAt  time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time       `db:"updated_at" json:"updated_at"`
	Title      string          `gorm:"type:varchar(255)" db:"title" json:"title" validate:"required,lte=255"`
	Slug       string          `gorm:"type:varchar(255)" db:"slug" json:"slug" validate:"required,lte=255"`
	ImageID    uint32          `gorm:"default:null" db:"image_id" json:"image_id"`
	Image      File            `json:"image"`
	Content    string          `db:"content" json:"content"`
	Votes      int32           `gorm:"default:0" db:"votes" json:"votes"`
	ProductID  uint32          `db:"product_id" json:"product_id" validate:"required"`
	Product    Product         `json:"product"`
	CategoryID uint32          `db:"category_id" json:"category_id" validate:"required"`
	Category   ArticleCategory `json:"category"`
	AuthorID   uint32          `db:"author_id" json:"author_id" validate:"required"`
	Author     User            `json:"author"`
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

type UserSavedArticle struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	ArticleID uint32    `gorm:"index:idx_article_user,unique" db:"article_id" json:"article_id"`
	Article   Article   `json:"article"`
	UserID    uint32    `gorm:"index:idx_article_user,unique" db:"user_id" json:"user_id"`
	User      Article   `json:"user"`
}
