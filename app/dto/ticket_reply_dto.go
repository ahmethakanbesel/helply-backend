package dto

type TicketReplyDTO struct {
	Content  string `json:"content"`
	TicketID uint32 `json:"ticket_id"`
}
