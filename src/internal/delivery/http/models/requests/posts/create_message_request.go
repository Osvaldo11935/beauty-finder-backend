package models_requests_posts

import "github.com/google/uuid"

type CreateMessageRequest struct {
	Type       int    `json:"type"`
	Body       string `json:"body"`
	SenderId   uuid.UUID `json:"senderId"`
	ReceiverId uuid.UUID `json:"receiverId"`
}