package primitives

import (
	"github.com/google/uuid"
)

type IEntity interface {
	GetID() uuid.UUID
	SetID(id uuid.UUID)
}
