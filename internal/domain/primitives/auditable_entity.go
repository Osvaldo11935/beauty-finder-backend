package primitives

import (
	"time"
)

// IAuditableEntity interface
type IAuditableEntity interface {
	IEntity
	GetIsActive() bool
	SetIsActive(isActive bool)
	GetIsDeleted() bool
	SetIsDeleted(isDeleted bool)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}
