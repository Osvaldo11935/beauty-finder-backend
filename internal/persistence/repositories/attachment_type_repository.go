package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type AttachmentTypeRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAttachmentTypeRepository(db *gorm.DB) interfaces_repositories.IAttachmentTypeRepository {
	return AttachmentRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}