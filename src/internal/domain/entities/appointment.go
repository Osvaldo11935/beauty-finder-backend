package entities

import (
	extensios "src/internal/domain/extensions"
	"src/internal/domain/object_values"
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Appointment struct {
	primitives.BaseAuditableEntity
	ProviderId uuid.UUID `gorm:"column:ProviderId;"`
	ClientId uuid.UUID `gorm:"column:ClientId;"`
    ServiceId uuid.UUID	 `gorm:"column:ServiceId;"`
	StatusId uuid.UUID `gorm:"column:StatusId;"`
	StartDate extensios.CustomDate `gorm:"column:startDate;"`
	EndDate extensios.CustomDate `gorm:"column:endDate;"`
	Status *AppointmentStatus 
	Service *Service
	Client *User
	Provider *User
}

func(s *Appointment) TableName() string{
	return "Appointment"
}

func NewAppointment(providerIdId uuid.UUID, clientId uuid.UUID, serviceId uuid.UUID, startDate extensios.CustomDate, 
	endDate extensios.CustomDate) Appointment{

	return Appointment{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		ClientId: clientId,
		ServiceId: serviceId,
        StatusId:  object_values.STATUS_PENDING_ID,
		StartDate: startDate,
		EndDate: endDate,
		ProviderId:  providerIdId,
	}
}

func (s *Appointment) Update(serviceId uuid.UUID){
      s.ServiceId = serviceId
}

func(s *Appointment) UpdateStatus(statusId uuid.UUID){
	 s.StatusId = statusId
}