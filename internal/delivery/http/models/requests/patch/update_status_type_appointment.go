package models_requests_patchs

import "github.com/google/uuid"

type UpdateStatusTypeAppointmentRequest struct {
	StatusTypeId *uuid.UUID `json:"statusTypeId"`
	Reason       *string    `json:"reason"`
}
