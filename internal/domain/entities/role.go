package entities

import "src/internal/domain/primitives"

type Role struct {
	primitives.BaseAuditableEntity
	Name string `gorm:"column:Name;" json:"name"`
	Users []*User `gorm:"foreignKey:RoleId;references:Id" json:"users"`
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