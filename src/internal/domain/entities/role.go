package entities

import "src/internal/domain/primitives"

type Role struct {
	primitives.BaseAuditableEntity
	Name string `gorm:"column:Name;"`
	Users []*User `gorm:"foreignKey:RoleId;references:ID"`
}

func(s *Role) TableName() string{
	return "Role"
}


func NewRole(name string) Role{
	body := Role{
	   BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
	   Name: name,
   }

   return body
}

func(s *Role) Update(name string){
	 s.Name = name
}