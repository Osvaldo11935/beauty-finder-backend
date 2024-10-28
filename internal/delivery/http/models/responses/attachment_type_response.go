package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type AttachmentTypeResponse struct {
	Id          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
}

func ToAttachmentTypeResponse(data *entities.AttachmentType) *AttachmentTypeResponse {
	if data == nil {
		return nil
	}
	return &AttachmentTypeResponse{
		Id:          data.ID,
		Type:        data.Type,
		Description: data.Description,
	}
}

func ToListAttachmentTypeResponse(data []entities.AttachmentType) []AttachmentTypeResponse {
	var resp []AttachmentTypeResponse

	for _, item := range data {
		resp = append(resp, *ToAttachmentTypeResponse(&item))
	}

	return resp
}
