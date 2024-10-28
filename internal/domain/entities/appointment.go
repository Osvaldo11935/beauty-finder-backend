package entities

import (
	extensios "src/internal/domain/extensions"
	"src/internal/domain/object_values"
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Appointment struct {
	primitives.BaseAuditableEntity
	ProviderId *uuid.UUID           `gorm:"column:ProviderId;" json:"providerId"`
	ClientId   uuid.UUID            `gorm:"column:ClientId;" json:"clientId"`
	ServiceId  uuid.UUID            `gorm:"column:ServiceId;" json:"serviceId"`
	StatusId   uuid.UUID            `gorm:"column:StatusId;" json:"statusId"`
	StartDate  extensios.CustomDate `gorm:"column:StartDate;" json:"startDate"`
	EndDate    extensios.CustomDate `gorm:"column:EndDate;" json:"endDate"`
	Address    *Address             `gorm:"foreignKey:AppointmentId;references:Id" json:"address"`
	Status     *AppointmentStatus   `json:"status"`
	Service    *Service             `json:"service"`
	Client     *User                `json:"client"`
	Provider   *User                `json:"provider"`
}

func (s *Appointment) TableName() string {
	return "Appointment"
}

func NewAppointment(providerIdId *uuid.UUID, clientId uuid.UUID, serviceId uuid.UUID, startDate extensios.CustomDate,
	endDate extensios.CustomDate) Appointment {

	return Appointment{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		ClientId:            clientId,
		ServiceId:           serviceId,
		StatusId:            object_values.STATUS_PENDING_ID,
		StartDate:           startDate,
		EndDate:             endDate,
		ProviderId:          providerIdId,
	}
}

func (s *Appointment) Update(serviceId uuid.UUID) {
	s.ServiceId = serviceId
}
func (s *Appointment) SetProvider(providerId uuid.UUID) {
	s.ProviderId = &providerId
}
func (s *Appointment) UpdateStatus(statusId uuid.UUID) {
	s.StatusId = statusId
}
