package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Attachment struct {
	primitives.BaseAuditableEntity
	Url              string     `gorm:"column:Url;"`
	UserId           *uuid.UUID `gorm:"column:UserId;"`
	ServiceId        *uuid.UUID `gorm:"column:ServiceId;"`
	CategoryId       *uuid.UUID `gorm:"column:CategoryId;"`
	AttachmentTypeId string     `gorm:"column:AttachmentTypeId;"`
	AttachmentType   *AttachmentType
}

func (s *Attachment) TableName() string {
	return "Attachment"
}
func NewAttachmentCategory(url string, categoryId uuid.UUID) Attachment {

	return Attachment{
		Url:       url,
		CategoryId: &categoryId,
	}
}
func NewAttachmentService(url string, serviceId uuid.UUID) Attachment {

	return Attachment{
		Url:       url,
		ServiceId: &serviceId,
	}
}
func NewAttachmentUser(url string, userId uuid.UUID) Attachment {

	return Attachment{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Url:    url,
		UserId: &userId,
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
