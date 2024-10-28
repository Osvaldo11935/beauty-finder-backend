package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Service struct {
	primitives.BaseAuditableEntity
	Name            string             `gorm:"column:Name" json:"name"`
	Description     string             `gorm:"column:Description" json:"description"`
	CategoryId      uuid.UUID          `gorm:"column:CategoryId" json:"categoryId"`
	Appointments    []*Appointment     `gorm:"foreignKey:ServiceId;references:Id" json:"appointments"`
	Price           *ServicePrice      `gorm:"foreignKey:ServiceId;references:Id" json:"price"`
	ServiceProvider *[]ServiceProvider `gorm:"foreignKey:ServiceId;references:Id" json:"serviceProvider"`
	Attachment      *Attachment        `gorm:"foreignKey:ServiceId;references:Id" json:"attachment"`
	Category        *ServiceCategory
}

func (s *Service) TableName() string {
	return "Service"
}

func NewService(name string, description string, categoryId uuid.UUID) Service {
	body := Service{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Name:                name,
		Description:         description,
		CategoryId:          categoryId,
	}

	return body
}

func (s *Service) Update(name *string, description *string, categoryId *uuid.UUID) {
	if name != nil {
		s.Name = *name
	}
	if description != nil {
		s.Description = *description
	}
	if categoryId != nil {
		s.CategoryId = *categoryId
	}
}
