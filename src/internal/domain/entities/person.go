package entities

import (
	extensios "src/internal/domain/extensions"
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Person struct {
	primitives.BaseAuditableEntity
	FullName              string               `gorm:"column:FullName;"`
	BirthDate             extensios.CustomDate `gorm:"column:BirthDate;"`
	Gender                string               `gorm:"column:Gender;"`
	Naturalness           string               `gorm:"column:Naturalness;"`
	MaritalStatus         string               `gorm:"column:MaritalStatus;"`
	FatherName            string               `gorm:"column:FatherName;"`
	MotherName            string               `gorm:"column:MotherName;"`
	NationalRegistry      string               `gorm:"column:NationalRegistry;"`
	PlaceIssuanceDocument string               `gorm:"column:PlaceIssuanceDocument;"`
	DateIssueDocument     extensios.CustomDate `gorm:"column:DateIssueDocument;"`
	UserId                *uuid.UUID           `gorm:"column:UserId;"`
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
