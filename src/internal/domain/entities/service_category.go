package entities

import "src/internal/domain/primitives"

type ServiceCategory struct {
	primitives.BaseAuditableEntity
	Name        string `gorm:"column:Name;"`
	Description string `gorm:"column:Description;"`
	Services []*Service `gorm:"foreignKey:CategoryId;references:ID"`
	Attachment *Attachment `gorm:"foreignKey:CategoryId;references:ID"`
}

func(s *ServiceCategory) TableName() string{
	return "ServiceCategory"
}

func NewServiceCategory(name string, description string) ServiceCategory{
	body := ServiceCategory{
	   BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
	   Name: name,
	   Description: description,
   }

   return body
}

func(s *ServiceCategory) Update(name *string, description *string){
	if name != nil {
	   s.Name = *name
	}
	if description != nil{
       s.Description = *description
	}
}