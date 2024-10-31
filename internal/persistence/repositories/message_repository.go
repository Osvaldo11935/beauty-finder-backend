package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type MessageRepository struct {
	*repositories_common.GormBaseRepository
}

func NewMessageRepository(db *gorm.DB) interfaces_repositories.IMessageRepository {
	return MessageRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}