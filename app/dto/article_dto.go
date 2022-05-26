package dto

type ArticleDTO struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	ProductID  uint32 `json:"product_id"`
	CategoryID uint32 `json:"category_id"`
	ImageID    uint32 `json:"image_id"`
}
