package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type AttachmentRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAttachmentRepository(db *gorm.DB) interfaces_repositories.IAttachmentRepository{
	return AttachmentRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}