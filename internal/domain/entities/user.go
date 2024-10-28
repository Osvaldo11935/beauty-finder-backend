package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	primitives.BaseAuditableEntity
	Email               *string            `gorm:"column:Email" json:"email"`
	UserName            *string            `gorm:"column:UserName" json:"userName"`
	Password            *string            `gorm:"column:Password" json:"password"`
	PhoneNumber         string             `gorm:"column:PhoneNumber" json:"phoneNumber"`
	RoleId              uuid.UUID          `gorm:"column:RoleId" json:"roleId"`
	Person              *Person            `gorm:"foreignKey:UserId;references:ID" json:"person"`
	ServicesProvided    []*ServiceProvider `gorm:"foreignKey:ProviderId;references:ID" json:"servicesProvided"`
	Conn                *websocket.Conn    `gorm:"-"`
	Attachment          []*Attachment      `gorm:"foreignKey:UserId;references:Id" json:"attachment"`
	MessagesReceiver    []*Message         `gorm:"foreignKey:ReceiverId;references:Id" json:"messagesReceiver"`
	MessagesSender      []*Message         `gorm:"foreignKey:SenderId;references:Id" json:"messagesSender"`
	AppointmentProvider []*Appointment     `gorm:"foreignKey:ProviderId;references:Id" json:"appointmentProvider"`
	AppointmentClient   []*Appointment     `gorm:"foreignKey:ClientId;references:Id" json:"appointmentClient"`
	Address             *Address           `gorm:"foreignKey:UserId;references:Id" json:"address"`
	FcmToken            []*FcmToken        `gorm:"foreignKey:UserId;references:Id" json:"fcmToken"`
	UserEvaluator       *UserRating        `gorm:"foreignKey:UserEvaluatorId;references:Id" json:"userEvaluator"`
	UserAvaluated       *UserRating        `gorm:"foreignKey:UserAvaluatedId;references:Id" json:"userAvaluated"`
	Role                *Role
}

func (s *User) TableName() string {
	return "User"
}

func NewUser(email *string, userName *string, password *string, phoneNumber string, roleId uuid.UUID) User {
	body := User{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Email:               email,
		UserName:            userName,
		Password:            password,
		PhoneNumber:         phoneNumber,
		RoleId:              roleId,
	}

	return body
}

func (s *User) Update(email *string, userName *string, password *string, phoneNumber *string, roleId *uuid.UUID) {
	if email != nil {
		s.Email = email
	}
	if userName != nil {
		s.UserName = userName
	}
	if password != nil {
		s.Password = password
	}
	if phoneNumber != nil {
		s.PhoneNumber = *phoneNumber
	}
	if roleId != nil {
		s.RoleId = *roleId
	}
}

func (s *User) SetServicesProvided(serviceIds []uuid.UUID) {
	for _, item := range serviceIds {
		req := ServiceProvider{
			ServiceId:  item,
			ProviderId: s.ID,
		}
		s.ServicesProvided = append(s.ServicesProvided, &req)
	}
}
func (s *User) SetFcmToken(fcmToken string, deviceName string, deviceId string) {
	req := FcmToken{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		TokenFcm:            fcmToken,
		DeviceName:          deviceName,
		DeviceId:            deviceId,
		UserId:              s.ID,
	}
	s.FcmToken = append(s.FcmToken, &req)
}

func (s *User) SendMessage(message *Message) error {
	if s.Conn != nil {
		return s.Conn.WriteJSON(message)
	}
	return nil
}

func (s *User) CloseConnection() {
	if s.Conn != nil {
		s.Conn.Close()
	}
}
