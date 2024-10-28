package models_responses

import (
	"src/internal/domain/entities"

	"github.com/google/uuid"
)

type AppointmentStatusResponse struct {
	Id          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
}

func ToAppointmentStatusResponse(data *entities.AppointmentStatus) *AppointmentStatusResponse {
	if data == nil {
		return nil
	}
	return &AppointmentStatusResponse{
		Id:          data.ID,
		Type:        data.Type,
		Description: data.Description,
	}
}

func ToListAppointmentStatusResponse(data []entities.AppointmentStatus) []AppointmentStatusResponse {
	var resp []AppointmentStatusResponse

	for _, item := range data {
		resp = append(resp, *ToAppointmentStatusResponse(&item))
	}

	return resp
}
