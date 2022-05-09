package models

import (
	"time"
)

type Ticket struct {
	ID             uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	CustomerID     uint32    `db:"customer_id" json:"customer_id" validate:"required"`
	Customer       User
	AgentID        uint32 `db:"agent_id" json:"agent_id" validate:"uuid"`
	Agent          User
	TicketTopicID  uint32 `db:"topic_id" json:"topic_id" validate:"required"`
	TicketTopic    TicketTopic
	TicketStatusID uint32 `gorm:"default:0" db:"status_id" json:"status_id"`
	TicketStatus   TicketStatus
	ProductID      uint32 `db:"product_id" json:"product_id" validate:"required"`
	Product        Product
}

type TicketStatus struct {
	ID               uint32 `db:"id" json:"id" validate:"required"`
	Name             string `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,lte=255"`
	AllowReply       bool   `gorm:"default:true" db:"allow_reply" json:"allow_reply"`
	HideFromCustomer bool   `gorm:"default:false" db:"hide_from_customer" json:"hide_from_customer"`
}

type TicketTopic struct {
	ID   uint32 `db:"id" json:"id" validate:"required"`
	Name string `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,lte=255"`
}

type TicketReply struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	TicketID  uint32    `db:"ticket_id" json:"ticket_id" validate:"required"`
	Ticket    Ticket
	UserID    uint32 `db:"user_id" json:"user_id" validate:"required"`
	User      User
	Content   string `db:"content" json:"content" validate:"required"`
}

type TicketReplyAttachment struct {
	ID            uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	TicketReplyID uint32    `db:"ticket_reply_id" json:"ticket_reply_id"`
	TicketReply   TicketReply
	FileID        uint32 `db:"file_id" json:"file_id"`
	File          File
}
