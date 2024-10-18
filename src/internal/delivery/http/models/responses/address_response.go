package models_responses

import (
	"src/internal/domain/entities"
	"strings"

	"github.com/google/uuid"
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
	Address      string    `json:"address"`
}

func GenerateAddress(data *entities.Address) string {
	var parts []string

	if data.Country != "" {
		parts = append(parts, data.Country)
	}
	if data.Province != "" {
		parts = append(parts, data.Province)
	}
	if data.City != "" {
		parts = append(parts, data.City)
	}
	if data.Commune != "" {
		parts = append(parts, data.Commune)
	}
	if data.District != "" {
		parts = append(parts, data.District)
	}
	if data.Neighborhood != "" {
		parts = append(parts, data.Neighborhood)
	}
	if data.Street != "" {
		parts = append(parts, data.Street)
	}
	return strings.Join(parts, ",")
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
		Address:      GenerateAddress(data),
	}
}

func ToListAddressResponse(data []entities.Address) []AddressResponse {
	var resp []AddressResponse

	for _, item := range data {
		resp = append(resp, *ToAddressResponse(&item))
	}

	return resp
}
