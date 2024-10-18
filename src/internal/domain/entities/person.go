package entities

import (
	extensios "src/internal/domain/extensions"
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Person struct {
	primitives.BaseAuditableEntity
	FullName              string               `gorm:"column:FullName;" json:"fullName"`
	BirthDate             extensios.CustomDate `gorm:"column:BirthDate;" json:"birthDate"`
	Gender                string               `gorm:"column:Gender;" json:"gender"`
	Naturalness           string               `gorm:"column:Naturalness;" json:"naturalness"`
	MaritalStatus         string               `gorm:"column:MaritalStatus;" json:"maritalStatus"`
	FatherName            string               `gorm:"column:FatherName;" json:"fatherName"`
	MotherName            string               `gorm:"column:MotherName;" json:"motherName"`
	NationalRegistry      string               `gorm:"column:NationalRegistry;" json:"nationalRegistry"`
	PlaceIssuanceDocument string               `gorm:"column:PlaceIssuanceDocument;" json:"placeIssuanceDocument"`
	DateIssueDocument     extensios.CustomDate `gorm:"column:DateIssueDocument;" json:"dateIssueDocument"`
	UserId                *uuid.UUID           `gorm:"column:UserId;" json:"userId"`
}

func (s *Person) TableName() string {
	return "Person"
}

func NewPerson(fullName string, birthDate extensios.CustomDate, gender string, naturalness string, maritalStatus string,
	fatherName string, motherName string, nationalRegistry string, placeIssuanceDocument string, dateIssueDocument extensios.CustomDate, userId *uuid.UUID) Person {
	return Person{
		BaseAuditableEntity:   *primitives.NewBaseAuditableEntity(),
		FullName:              fullName,
		BirthDate:             birthDate,
		Gender:                gender,
		Naturalness:           naturalness,
		MaritalStatus:         maritalStatus,
		FatherName:            fatherName,
		MotherName:            motherName,
		NationalRegistry:      nationalRegistry,
		PlaceIssuanceDocument: placeIssuanceDocument,
		DateIssueDocument:     dateIssueDocument,
		UserId:                userId,
	}
}

func (s *Person) Update(fullName *string, birthDate *extensios.CustomDate, gender *string, naturalness *string, maritalStatus *string,
	fatherName *string, motherName *string, nationalRegistry *string, placeIssuanceDocument *string, dateIssueDocument *extensios.CustomDate, userId *uuid.UUID) {
	if fullName != nil {
		s.FullName = *fullName
	}

	if birthDate != nil {
		s.BirthDate = *birthDate
	}

	if gender != nil {
		s.Gender = *gender
	}
	if naturalness != nil {
		s.Naturalness = *naturalness
	}
	if maritalStatus != nil {
		s.MaritalStatus = *maritalStatus
	}
	if fatherName != nil {
		s.FatherName = *fatherName
	}
	if motherName != nil {
		s.MotherName = *motherName
	}
	if nationalRegistry != nil {
		s.NationalRegistry = *nationalRegistry
	}
	if placeIssuanceDocument != nil {
		s.PlaceIssuanceDocument = *placeIssuanceDocument
	}
	if dateIssueDocument != nil {
		s.DateIssueDocument = *dateIssueDocument
	}
	if userId != nil {
		s.UserId = userId
	}
}
