package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type RatingTypeRepository struct {
	*repositories_common.GormBaseRepository
}

func NewRatingTypeRepository(db *gorm.DB) interfaces_repositories.IRatingTypeRepository {
	return RatingTypeRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}
