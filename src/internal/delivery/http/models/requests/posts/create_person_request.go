package models_requests_posts

import (
	extensios "src/internal/domain/extensions"
)

type CreatePersonRequest struct{
	FullName string `json:"fullName"`
	BirthDate extensios.CustomDate `json:"birthDate"`
	Gender                string `json:"gender"`              
	Naturalness           string `json:"naturalness"`            
	MaritalStatus         string `json:"maritalStatus"`           
	FatherName            string `json:"fatherName"`              
	MotherName            string `json:"motherName"`                         
	NationalRegistry      string `json:"nationalRegistry"`              
	PlaceIssuanceDocument string `json:"placeIssuanceDocument"`              
	DateIssueDocument     extensios.CustomDate `json:"dateIssueDocument"`
}