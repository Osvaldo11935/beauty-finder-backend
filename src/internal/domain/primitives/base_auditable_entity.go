package primitives

import (
	"time"
)

type BaseAuditableEntity struct {
	BaseEntity
	IsActive  bool      `gorm:"column:IsActive;" json:"isActive"`
	IsDeleted bool      `gorm:"column:IsDeleted;" json:"isDeleted"`
	CreatedAt time.Time `gorm:"column:CreatedAt;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;" json:"updatedAt"`
}

func NewBaseAuditableEntity() *BaseAuditableEntity {
	return &BaseAuditableEntity{
		BaseEntity: *NewBaseEntity(),
		IsActive:   true,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (bae *BaseAuditableEntity) GetIsActive() bool {
	return bae.IsActive
}

func (bae *BaseAuditableEntity) SetIsActive(isActive bool) {
	bae.IsActive = isActive
}

func (bae *BaseAuditableEntity) GetIsDeleted() bool {
	return bae.IsDeleted
}

func (bae *BaseAuditableEntity) SetIsDeleted(isDeleted bool) {
	bae.IsDeleted = isDeleted
}

func (bae *BaseAuditableEntity) GetCreatedAt() time.Time {
	return bae.CreatedAt
}

func (bae *BaseAuditableEntity) SetCreatedAt(createdAt time.Time) {
	bae.CreatedAt = createdAt
}

func (bae *BaseAuditableEntity) GetUpdatedAt() time.Time {
	return bae.UpdatedAt
}

func (bae *BaseAuditableEntity) SetUpdatedAt(updatedAt time.Time) {
	bae.UpdatedAt = updatedAt
}
