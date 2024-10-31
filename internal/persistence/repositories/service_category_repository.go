package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type ServiceCategoryRepository struct {
	*repositories_common.GormBaseRepository
}

func NewServiceCategoryRepository(db *gorm.DB) interfaces_repositories.IServiceCategoryRepository {
	return ServiceCategoryRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}