package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type UserRatingRepository struct {
	*repositories_common.GormBaseRepository
}

func NewUserRatingRepository(db *gorm.DB) interfaces_repositories.IUserRatingRepository {
	return UserRatingRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}