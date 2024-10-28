package models_requests_puts

import "github.com/google/uuid"

type UpdateMessageRequest struct {
	Type       *int     `json:"type"`
	Body       *string  `json:"body"`
	SenderId   *uuid.UUID `json:"senderId"`
	ReceiverId *uuid.UUID  `json:"receiverId"`
}