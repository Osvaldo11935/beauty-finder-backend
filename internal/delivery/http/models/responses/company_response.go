package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type CompanyResponse struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	PhoneNumber      string    `json:"phoneNumber"`
	NationalRegistry string    `json:"nationalRegistry"`
}

func ToCompanyResponse(data *entities.Company) *CompanyResponse {
	if data == nil {
		return nil
	}
	return &CompanyResponse{
		Id:               data.ID,
		Name:             data.Name,
		PhoneNumber:      data.PhoneNumber,
		NationalRegistry: data.NationalRegistry,
	}
}

func ToListCompanyResponse(data []*entities.Company) []CompanyResponse {
	var resp []CompanyResponse

	for _, item := range data {
		resp = append(resp, *ToCompanyResponse(item))
	}

	return resp
}
