package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type ServiceResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price    *ServicePriceResponse `json:"price"`
	Category    *ServiceCategoryResponse `json:"category"`
}

func ToServiceResponse(data *entities.Service) *ServiceResponse {
	if data == nil {
		return nil
	 }
	return &ServiceResponse{
		Id:          data.ID,
		Name: data.Name,
		Description: data.Description,
		Category: ToServiceCategoryResponse(data.Category),
		Price: ToServicePriceResponse(data.Price),
	}
}

func ToListServiceResponse(data []entities.Service) []ServiceResponse {
	var resp []ServiceResponse

	for _, item := range data {
		resp = append(resp, *ToServiceResponse(&item))
	}

	return resp
}