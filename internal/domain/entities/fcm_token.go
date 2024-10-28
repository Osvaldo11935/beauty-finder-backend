package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type FcmToken struct {
	primitives.BaseAuditableEntity
	UserId uuid.UUID `gorm:"column:UserId;" json:"userId" `
	TokenFcm string `gorm:"column:TokenFcm;" json:"tokenFcm"`
	DeviceId string `gorm:"column:DeviceId;" json:"deviceId"`
	DeviceName string `gorm:"column:DeviceName;" json:"deviceName"`
}


func (s *FcmToken) TableName() string {
	return "FcmToken"
}