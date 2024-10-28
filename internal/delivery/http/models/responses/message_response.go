package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Id         uuid.UUID     `json:"id"`
	Type       int           `json:"type"`
	Body       string        `json:"body"`
	Sender     *UserResponse `json:"sender"`
	Receiver   *UserResponse `json:"receiver"`
	SenderId   uuid.UUID     `json:"senderId"`
	ReceiverId uuid.UUID     `json:"receiverId"`
}

func ToMessageResponse(data *entities.Message) *MessageResponse {
	if data == nil {
		return nil
	}
	d := &MessageResponse{
		Id:       data.ID,
		Type:     data.Type,
		Body:     data.Body,
		Sender:   ToUserResponse(data.Sender),
		Receiver: ToUserResponse(data.Receiver),
	}
	if data.Sender != nil {
		d.SenderId = data.SenderId
	}
	if data.Receiver != nil {
		d.ReceiverId = data.ReceiverId
	}
	return d
}

func ToListMessageResponse(data []entities.Message) []MessageResponse {
	var resp []MessageResponse

	for _, item := range data {
		resp = append(resp, *ToMessageResponse(&item))
	}

	return resp
}
