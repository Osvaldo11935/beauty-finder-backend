package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type MessageRepository struct {
	*repositories_common.GormBaseRepository
}

func NewMessageRepository() interfaces_repositories.IMessageRepository {
	return MessageRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}