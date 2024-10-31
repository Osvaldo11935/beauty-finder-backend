package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type UserRepository struct {
	*repositories_common.GormBaseRepository
}

func NewUserRepository(db *gorm.DB) interfaces_repositories.IUserRepository {
	return UserRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}