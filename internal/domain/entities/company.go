package entities

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/primitives"
)

type Company struct {
	primitives.BaseAuditableEntity
	Name             string        `gorm:"column:Type;" json:"name"`
	PhoneNumber      string        `gorm:"column:PhoneNumber;" json:"phoneNumber"`
	NationalRegistry string        `gorm:"column:NationalRegistry;" json:"nationalRegistry"`
	EmployeeNumber   int           `gorm:"column:EmployeeNumber" json:"employeeNumber"`
	Address          *Address      `gorm:"foreignKey:CompanyId;references:Id" json:"address"`
	Attachment       []*Attachment `gorm:"foreignKey:CompanyId;references:Id" json:"attachment"`
	User             *User         `gorm:"foreignKey:CompanyId;references:Id" json:"User"`
}

func (s *Company) TableName() string {
	return "Company"
}

func NewCompany(request models_requests_posts.CreateCompanyRequest) Company {

	return Company{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Name:                request.Name,
		PhoneNumber:         request.PhoneNumber,
		NationalRegistry:    request.NationalRegistry,
		EmployeeNumber:      request.EmployeeNumber,
	}
}

func (s *Company) Update(request models_requests_puts.UpdateCompanyRequest) {
	if request.Name != nil {
		s.Name = *request.Name
	}

	if request.PhoneNumber != nil {
		s.PhoneNumber = *request.PhoneNumber
	}

	if request.NationalRegistry != nil {
		s.NationalRegistry = *request.NationalRegistry
	}

	if request.NationalRegistry != nil {
		s.NationalRegistry = *request.NationalRegistry
	}
}
