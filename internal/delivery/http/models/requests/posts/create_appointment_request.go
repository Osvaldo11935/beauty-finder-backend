package models_requests_posts

import (
	extensios "src/internal/domain/extensions"

	"github.com/google/uuid"
)

type CreateAppointmentRequest struct {
	ProviderId    *uuid.UUID `json:"providerId"`
	ClientId  uuid.UUID `json:"clientId"`
	ServiceId uuid.UUID `json:"serviceId"`
	StartDate extensios.CustomDate `json:"startDate"`
	EndDate extensios.CustomDate `json:"endDate"`
}