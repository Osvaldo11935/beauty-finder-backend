package entities

import "src/internal/domain/primitives"

type AppointmentStatus struct {
	primitives.BaseAuditableEntity
	Type string `gorm:"column:Type;" json:"type"`
	Description string `gorm:"column:Description;" json:"description"`
	Appointments []*Appointment `gorm:"foreignKey:StatusId;references:Id" json:"appointments"`
}


func(s *AppointmentStatus) TableName() string {
	return "AppointmentStatus"
}

func NewAppointmentStatus(_type string, description string) AppointmentStatus{
	 body := AppointmentStatus{
        BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Type: _type,
		Description: description,
	}

	return body
}

func(s *AppointmentStatus) Update(_type *string, description *string){
     if _type != nil {
        s.Type = *_type
	 }
	 if description != nil {
        s.Description = *description
	 }
}