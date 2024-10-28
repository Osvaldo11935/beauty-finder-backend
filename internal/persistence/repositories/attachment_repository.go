package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type AttachmentRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAttachmentRepository() interfaces_repositories.IAttachmentRepository{
	return AttachmentRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}