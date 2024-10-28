package primitives

import (
	"github.com/google/uuid"
)

type BaseEntity struct {
	ID           uuid.UUID `gorm:"column:Id;primaryKey" json:"id"`
	domainEvents []interface{}
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		ID:           uuid.New(),
		domainEvents: []interface{}{},
	}
}

func (be *BaseEntity) GetID() uuid.UUID {
	return be.ID
}

func (be *BaseEntity) SetID(id uuid.UUID) {
	be.ID = id
}

func (be *BaseEntity) GetDomainEvents() []interface{} {
	return be.domainEvents
}

func (be *BaseEntity) AddDomainEvent(event interface{}) {
	be.domainEvents = append(be.domainEvents, event)
}

func (be *BaseEntity) RemoveDomainEvent(event BaseEvent) {
	for i, e := range be.domainEvents {
		if e == event {
			be.domainEvents = append(be.domainEvents[:i], be.domainEvents[i+1:]...)
			break
		}
	}
}

func (be *BaseEntity) ClearDomainEvents() {
	be.domainEvents = []interface{}{}
}
