package models_requests_puts

import "github.com/google/uuid"

type UpdateAppointmentRequest struct {
	UserId    uuid.UUID `json:"userId"`
	ClientId  uuid.UUID `json:"clientId"`
	ServiceId uuid.UUID `json:"serviceId"`
}