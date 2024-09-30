package models_requests_puts

import "github.com/google/uuid"

type UpdateServiceRequest struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	CategoryId  *uuid.UUID `json:"categoryId"`
}