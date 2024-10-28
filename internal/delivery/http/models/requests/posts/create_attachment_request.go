package models_requests_posts

import "github.com/google/uuid"

type CreateAttachmentRequest struct {
	Url              string     `json:"url"`
	UserId           *uuid.UUID `json:"userId"`
	ServiceId        *uuid.UUID `json:"serviceId"`
	CategoryId       *uuid.UUID `json:"categoryId"`
	AttachmentTypeId uuid.UUID  `json:"attachmentTypeId"`
}
