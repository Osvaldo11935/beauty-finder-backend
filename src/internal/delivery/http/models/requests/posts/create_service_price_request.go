package models_requests_posts

import "github.com/google/uuid"

type CreateServicePriceRequest struct {
	Amount    float64 `json:"amount"`
	ServiceId uuid.UUID `json:"serviceId"`
}