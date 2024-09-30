package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type AttachmentResponse struct {
	Id             uuid.UUID              `json:"id"`
	Url            string                 `json:"url"`
	AttachmentType *AttachmentTypeResponse `json:"attachmentType"`
}

func ToAttachmentResponse(data *entities.Attachment) *AttachmentResponse {
	if data == nil {
		return nil
	}
	return &AttachmentResponse{
		Id:             data.ID,
		Url:            data.Url,
		AttachmentType: ToAttachmentTypeResponse(data.AttachmentType),
	}
}

func ToListAttachmentResponse(data []entities.Attachment) []AttachmentResponse {
	var resp []AttachmentResponse

	for _, item := range data {
		resp = append(resp, *ToAttachmentResponse(&item))
	}

	return resp
}
