package models_requests_puts

import "github.com/google/uuid"

type UpdateServicePriceRequest struct {
	Amount    *float64   `json:"amount"`
	ServiceId *uuid.UUID `json:"serviceId"`
}