package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	primitives.BaseAuditableEntity
	Email               string             `gorm:"column:Email"`
	UserName            string             `gorm:"column:UserName"`
	Password            string             `gorm:"column:Password"`
	PhoneNumber         string             `gorm:"column:PhoneNumber"`
	RoleId              uuid.UUID          `gorm:"column:RoleId"`
	Person              *Person            `gorm:"foreignKey:UserId;references:ID"`
	ServicesProvided    []*ServiceProvider `gorm:"foreignKey:ProviderId;references:ID"`
	Conn                *websocket.Conn    `gorm:"-"`
	Attachment          []*Attachment      `gorm:"foreignKey:UserId;references:ID"`
	MessagesReceiver    []*Message         `gorm:"foreignKey:ReceiverId;references:ID"`
	MessagesSender      []*Message         `gorm:"foreignKey:SenderId;references:ID"`
	AppointmentProvider []*Appointment     `gorm:"foreignKey:ProviderId;references:ID"`
	AppointmentClient   []*Appointment     `gorm:"foreignKey:ClientId;references:ID"`
	Role                *Role
}

func (s *User) TableName() string {
	return "User"
}

func NewUser(email string, userName string, password string, phoneNumber string, roleId uuid.UUID) User {
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
		s.Email = *email
	}
	if userName != nil {
		s.UserName = *userName
	}
	if password != nil {
		s.Password = *password
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