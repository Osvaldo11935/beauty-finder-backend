package entities

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
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
	CompanyId           *uuid.UUID         `gorm:"column:CompanyId" json:"companyId"`
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
	Company             *Company
	Role                *Role
}

func (s *User) TableName() string {
	return "User"
}

func NewUser(roleId uuid.UUID, request models_requests_posts.CreateUserRequest) User {
	body := User{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Email:               request.Email,
		UserName:            request.UserName,
		Password:            request.Password,
		PhoneNumber:         request.PhoneNumber,
		RoleId:              roleId,
		CompanyId:           request.CompanyId,
	}

	return body
}

func (s *User) Update(request models_requests_puts.UpdateUserRequest) {
	if request.Email != nil {
		s.Email = request.Email
	}
	if request.UserName != nil {
		s.UserName = request.UserName
	}
	if request.Password != nil {
		s.Password = request.Password
	}
	if request.PhoneNumber != nil {
		s.PhoneNumber = *request.PhoneNumber
	}
	if request.RoleId != nil {
		s.RoleId = *request.RoleId
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
