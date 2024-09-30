package entities

import (
	"github.com/google/uuid"
)

type ServiceProvider struct {
	 ServiceId uuid.UUID `gorm:"column:ServiceId"`
	 ProviderId uuid.UUID `gorm:"column:ProviderId"`
	 Service *Service
	 Provider *User

}

func(s *ServiceProvider) TableName() string{
	return "ServiceProvider"
}
