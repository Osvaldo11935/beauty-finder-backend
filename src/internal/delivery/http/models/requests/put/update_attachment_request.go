package models_requests_puts

import "github.com/google/uuid"

type UpdateAttachmentRequest struct {
	Url              *string  `json:"url"`
	UserId           *uuid.UUID `json:"userId"`
	AttachmentTypeId *string  `json:"attachmentTypeId"`
}