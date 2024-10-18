package entities

import "src/internal/domain/primitives"

type ServiceCategory struct {
	primitives.BaseAuditableEntity
	Name        string `gorm:"column:Name;" json:"name"`
	Description string `gorm:"column:Description;" json:"description"`
	Services []*Service `gorm:"foreignKey:CategoryId;references:Id" json:"services"`
	Attachment *Attachment `gorm:"foreignKey:CategoryId;references:Id" json:"attachment"`
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