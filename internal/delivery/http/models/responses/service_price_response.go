package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type ServicePriceResponse struct {
	Id     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
}

func ToServicePriceResponse(data *entities.ServicePrice) *ServicePriceResponse {
	if data == nil {
		return nil
	}
	return &ServicePriceResponse{
		Id:     data.ID,
		Amount: data.Amount,
	}
}

func ToListServicePriceResponse(data []entities.ServicePrice) []ServicePriceResponse {
	var resp []ServicePriceResponse

	for _, item := range data {
		d := ToServicePriceResponse(&item)
		resp = append(resp, *d)
	}

	return resp
}
