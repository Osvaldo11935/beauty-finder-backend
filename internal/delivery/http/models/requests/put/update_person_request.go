package models_requests_puts

import (
	extensios "src/internal/domain/extensions"
	"github.com/google/uuid"
)

type UpdatePersonRequest struct{
	FullName *string `json:"type"`
	BirthDate *extensios.CustomDate `json:"birthDate"`
	Gender                *string `json:"gender"`              
	Naturalness           *string `json:"naturalness"`            
	MaritalStatus         *string `json:"maritalStatus"`           
	FatherName            *string `json:"fatherName"`              
	MotherName            *string `json:"motherName"`                         
	NationalRegistry      *string `json:"nationalRegistry"`              
	PlaceIssuanceDocument *string `json:"placeIssuanceDocument"`              
	DateIssueDocument     *extensios.CustomDate `json:"dateIssueDocument"`
	UserId *uuid.UUID `json:"userId"`
}