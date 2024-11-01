package entities

import (
	"src/internal/domain/primitives"
)

type AttachmentType struct {
	primitives.BaseAuditableEntity
	Type string `gorm:"column:Type;" json:"type"`
	Description string `gorm:"column:Description;" json:"description"`
	Attachment []*Attachment `gorm:"foreignKey:AttachmentTypeId;references:Id" json:"attachment"`
}


func(s *AttachmentType) TableName() string{
	return "AttachmentType"
}

func NewAttachmentType(_type string, description string) AttachmentType{

	return AttachmentType{
		   BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
           Type: _type,
		   Description: description,
	}
}

func (s *AttachmentType) Update(_type *string, description *string){
      if _type != nil {
		  s.Type = *_type
	  }

	  if description != nil {
		s.Description = *description
	}

}