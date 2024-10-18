package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type UserRatingRepository struct {
	*repositories_common.GormBaseRepository
}

func NewUserRatingRepository() interfaces_repositories.IUserRatingRepository {
	return UserRatingRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}