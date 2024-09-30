package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Attachment struct {
	primitives.BaseAuditableEntity
	Url string  `gorm:"column:Url;"`
	UserId uuid.UUID `gorm:"column:UserId;"`
	AttachmentTypeId string `gorm:"column:AttachmentTypeId;"`
	AttachmentType *AttachmentType
}


func(s *Attachment) TableName() string{
	return "Attachment"
}


func NewAttachment(url string, userId uuid.UUID) Attachment{

	return Attachment{ 
		Url: url,
		UserId: userId,
	}
}

func (s *Attachment) Update(url *string, userId *uuid.UUID){
      if url != nil {
		  s.Url = *url
	  }

	  if userId != nil {
		s.UserId = *userId
	}
}