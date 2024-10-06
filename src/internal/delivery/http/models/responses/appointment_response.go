package models_responses

import (
	"src/internal/domain/entities"
	extensios "src/internal/domain/extensions"

	"github.com/google/uuid"
)

type AppointmentResponse struct {
	Id        uuid.UUID                  `json:"id"`
	StartDate extensios.CustomDate       `json:"startDate"`
	EndDate   extensios.CustomDate       `json:"endDate"`
	Client    *UserResponse              `json:"client"`
	Provider  *UserResponse              `json:"provider"`
	Status    *AppointmentStatusResponse `json:"status"`
	Service   *ServiceResponse           `json:"service"`
}

func ToAppointmentResponse(data *entities.Appointment) *AppointmentResponse {
	if data == nil {
		return nil
	}
	return &AppointmentResponse{
		Id:        data.ID,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		Client:    ToUserResponse(data.Client),
		Provider:  ToUserResponse(data.Provider),
		Status:    ToAppointmentStatusResponse(data.Status),
		Service:   ToServiceResponse(data.Service),
	}
}

func ToListAppointmentResponse(data []entities.Appointment) []AppointmentResponse {
	var resp []AppointmentResponse

	for _, item := range data {
		resp = append(resp, *ToAppointmentResponse(&item))
	}

	return resp
}
