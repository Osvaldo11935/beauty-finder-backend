package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Attachment struct {
	primitives.BaseAuditableEntity
	Url              string     `gorm:"column:Url;" json:"url"`
	UserId           *uuid.UUID `gorm:"column:UserId;" json:"userId"`
	ServiceId        *uuid.UUID `gorm:"column:ServiceId;" json:"serviceId"`
	CategoryId       *uuid.UUID `gorm:"column:CategoryId;" json:"categoryId"`
	AttachmentTypeId uuid.UUID     `gorm:"column:AttachmentTypeId;" json:"attachmentTypeId"`
	AttachmentType   *AttachmentType
}

func (s *Attachment) TableName() string {
	return "Attachment"
}
func NewAttachmentCategory(url string, categoryId uuid.UUID, attachmentTypeId uuid.UUID) Attachment {

	return Attachment{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Url:        url,
		CategoryId: &categoryId,
		AttachmentTypeId: attachmentTypeId,
	}
}
func NewAttachmentService(url string, serviceId uuid.UUID, attachmentTypeId uuid.UUID) Attachment {

	return Attachment{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Url:       url,
		ServiceId: &serviceId,
		AttachmentTypeId: attachmentTypeId,
	}
}
func NewAttachmentUser(url string, userId uuid.UUID, attachmentTypeId uuid.UUID) Attachment {

	return Attachment{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Url:                 url,
		UserId:              &userId,
		AttachmentTypeId: attachmentTypeId,
	}
}

func (s *Attachment) Update(url *string, userId *uuid.UUID, serviceId *uuid.UUID, categoryId *uuid.UUID) {
	if url != nil {
		s.Url = *url
	}
	if userId != nil {
		s.UserId = userId
	}
	if serviceId != nil {
		s.ServiceId = serviceId
	}
	if categoryId != nil {
		s.CategoryId = categoryId
	}
}
