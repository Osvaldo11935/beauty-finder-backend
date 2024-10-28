package models_responses

import (
	"src/internal/domain/entities"
	extensios "src/internal/domain/extensions"

	"github.com/google/uuid"
)

type PersonResponse struct {
	Id                    uuid.UUID            `json:"id"`
	FullName              string               `json:"fullName"`
	BirthDate             *extensios.CustomDate `json:"birthDate"`
	Gender                string               `json:"gender"`
	Naturalness           string               `json:"naturalness"`
	MaritalStatus         string               `json:"maritalStatus"`
	FatherName            string               `json:"fatherName"`
	MotherName            string               `json:"motherName"`
	NationalRegistry      string               `json:"nationalRegistry"`
	PlaceIssuanceDocument string               `json:"placeIssuanceDocument"`
	DateIssueDocument     *extensios.CustomDate `json:"dateIssueDocument"`
}


func PersonDataFromGovernmentResponseToPersonResponse(data *PersonDataFromGovernmentResponse) *PersonResponse {
    
	if data == nil {
		return nil
	 }
	return &PersonResponse{
		Id: data.Id,               
		FullName: data.FullName, 
		BirthDate: data.BirthDate,            
		Gender: data.Gender,               
		Naturalness: data.Naturalness,         
		MaritalStatus: data.MaritalStatus,        
		FatherName: data.FatherName,           
		MotherName: data.MotherName,           
		NationalRegistry: data.NationalRegistry,     
		PlaceIssuanceDocument: data.PlaceIssuanceDocument,
		DateIssueDocument: data.DateIssueDocument,    
	}
}

func ToPersonResponse(data *entities.Person) *PersonResponse {
    
	if data == nil {
		return nil
	 }
	 

	return &PersonResponse{
		Id: data.ID,               
		FullName: data.FullName, 
		BirthDate: &data.BirthDate,            
		Gender: data.Gender,               
		Naturalness: data.Naturalness,         
		MaritalStatus: data.MaritalStatus,        
		FatherName: data.FatherName,           
		MotherName: data.MotherName,           
		NationalRegistry: data.NationalRegistry,     
		PlaceIssuanceDocument: data.PlaceIssuanceDocument,
		DateIssueDocument: data.DateIssueDocument,    
	}
}

func ToListPersonResponse(data []entities.Person) []PersonResponse {
	var resp []PersonResponse

	for _, item := range data {
		resp = append(resp, *ToPersonResponse(&item))
	}

	return resp
}