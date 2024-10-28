package models_requests_puts

import "github.com/google/uuid"

type UpdateAppointmentRequest struct {
	ProviderId    *uuid.UUID `json:"providerId"`
	ClientId  *uuid.UUID `json:"clientId"`
	ServiceId *uuid.UUID `json:"serviceId"`
}