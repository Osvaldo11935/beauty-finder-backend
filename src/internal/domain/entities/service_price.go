package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type ServicePrice struct {
	primitives.BaseAuditableEntity
	Amount      float64 `gorm:"column:Amount"`
    ServiceId uuid.UUID `gorm:"column:ServiceId"`
	Service *Service 
}

func(s *ServicePrice) TableName() string{
	return "ServicePrice"
}

func NewServicePrice(amount float64, serviceId uuid.UUID) ServicePrice{
	body := ServicePrice{
	   BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
	   Amount: amount,
	   ServiceId: serviceId,
   }

   return body
}

func(s *ServicePrice) Update(amount *float64, serviceId *uuid.UUID){
	if amount != nil {
	   s.Amount = *amount
	}
	if serviceId != nil{
       s.ServiceId = *serviceId
	}
}