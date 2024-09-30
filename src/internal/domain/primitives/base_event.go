package primitives

import (
	"time"
)

type BaseEvent struct {
	DateOccurred time.Time
}

func NewBaseEvent() *BaseEvent {
	return &BaseEvent{
		DateOccurred: time.Now().UTC(),
	}
}
