package entities

import (
	"github.com/google/uuid"
)

type ServiceProvider struct {
	ServiceId  uuid.UUID `gorm:"column:ServiceId" json:"serviceId"`
	ProviderId uuid.UUID `gorm:"column:ProviderId" json:"providerId"`
	Service    *Service  `json:"service"`
	Provider   *User     `json:"provider"`
}

func (s *ServiceProvider) TableName() string {
	return "ServiceProvider"
}
