package entities

import (
	"src/internal/domain/primitives"

	"github.com/google/uuid"
)

type Message struct {
	primitives.BaseAuditableEntity
	Type       int       `gorm:"column:Type;" json:"type"`
	Body       string    `gorm:"column:Body;" json:"body"`
	SenderId   uuid.UUID `gorm:"column:SenderId;" json:"senderId"`
	ReceiverId uuid.UUID `gorm:"column:ReceiverId;" json:"receiverId"`
	Receiver   *User `json:"receiver"`
	Sender     *User`json:"sender"`
}

func (s *Message) TableName() string {
	return "Message"
}

func NewMessage(_type int, body string, senderId uuid.UUID, receiverId uuid.UUID) Message {
	return Message{
		BaseAuditableEntity: *primitives.NewBaseAuditableEntity(),
		Type:                _type,
		Body:                body,
		SenderId:            senderId,
		ReceiverId:          receiverId,
	}
}

func (s *Message) Update(_type *int, body *string, senderId *uuid.UUID, receiverId *uuid.UUID) {
	if _type != nil {
		s.Type = *_type
	}

	if body != nil {
		s.Body = *body
	}

	if senderId != nil {
		s.SenderId = *senderId
	}
	if receiverId != nil {
		s.ReceiverId = *receiverId
	}
}
