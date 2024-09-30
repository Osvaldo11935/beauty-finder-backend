package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Id       uuid.UUID     `json:"id"`
	Type     int           `json:"type"`
	Body     string        `json:"body"`
	Receiver *UserResponse `json:"receiver"`
}

func ToMessageResponse(data *entities.Message) *MessageResponse {
	if data == nil {
		return nil
	}
	return &MessageResponse{
		Id:       data.ID,
		Type:     data.Type,
		Body:     data.Body,
		Receiver: ToUserResponse(data.Receiver),
	}
}

func ToListMessageResponse(data []entities.Message) []MessageResponse {
	var resp []MessageResponse

	for _, item := range data {
		resp = append(resp, *ToMessageResponse(&item))
	}

	return resp
}
