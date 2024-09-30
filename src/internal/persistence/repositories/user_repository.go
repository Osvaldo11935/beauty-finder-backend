package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type UserRepository struct {
	*repositories_common.GormBaseRepository
}

func NewUserRepository() interfaces_repositories.IUserRepository {
	return UserRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}