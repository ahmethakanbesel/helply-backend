package dto

type TicketDTO struct {
	TopicID   uint32 `json:"topic_id"`
	ProductID uint32 `json:"product_id"`
	Content   string `json:"content"`
}
