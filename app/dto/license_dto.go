package dto

import (
	"time"
)

type LicenseDTO struct {
	ExpiresAt time.Time `json:"expires_at"`
	Code      string    `json:"code"`
	ProductID uint32    `json:"product_id"`
	IsActive  bool      `json:"is_active"`
}
