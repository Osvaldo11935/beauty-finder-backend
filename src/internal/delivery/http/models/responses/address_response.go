package models_responses

import (
	"github.com/google/uuid"
	"src/internal/domain/entities"
)

type AddressResponse struct {
	Id           uuid.UUID `json:"id"`
	District     string    `json:"district"`
	Commune      string    `json:"commune"`
	Province     string    `json:"province"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Street       string    `json:"street"`
	Neighborhood string    `json:"neighborhood"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
}

func ToAddressResponse(data *entities.Address) *AddressResponse {

	if data == nil {
		return nil
	}

	return &AddressResponse{
		Id:           data.ID,
		District:     data.District,
		Commune:      data.Commune,
		Province:     data.Province,
		Country:      data.Country,
		City:         data.City,
		Street:       data.Street,
		Neighborhood: data.Neighborhood,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
	}
}

func ToListAddressResponse(data []entities.Address) []AddressResponse {
	var resp []AddressResponse

	for _, item := range data {
		resp = append(resp, *ToAddressResponse(&item))
	}

	return resp
}
