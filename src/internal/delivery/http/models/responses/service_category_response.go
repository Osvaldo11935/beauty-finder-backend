package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type ServiceCategoryResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func ToServiceCategoryResponse(data *entities.ServiceCategory) *ServiceCategoryResponse {
	if data == nil {
		return nil
	}
	return &ServiceCategoryResponse{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
}

func ToListServiceCategoryResponse(data []entities.ServiceCategory) []ServiceCategoryResponse {
	var resp []ServiceCategoryResponse

	for _, item := range data {
		resp = append(resp, *ToServiceCategoryResponse(&item))
	}

	return resp
}
