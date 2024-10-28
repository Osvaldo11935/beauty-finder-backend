package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type AttachmentTypeRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAttachmentTypeRepository() interfaces_repositories.IAttachmentTypeRepository {
	return AttachmentRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}