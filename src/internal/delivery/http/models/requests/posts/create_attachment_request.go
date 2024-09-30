package models_requests_posts

import "github.com/google/uuid"

type CreateAttachmentRequest struct {
	Url              string  `json:"url"`
	UserId           uuid.UUID `json:"userId"`
	AttachmentTypeId string  `json:"attachmentTypeId"`
}