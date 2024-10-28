package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type RatingTypeRepository struct {
	*repositories_common.GormBaseRepository
}

func NewRatingTypeRepository() interfaces_repositories.IRatingTypeRepository {
	return RatingTypeRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}
