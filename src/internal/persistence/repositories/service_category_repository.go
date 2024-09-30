package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type ServiceCategoryRepository struct {
	*repositories_common.GormBaseRepository
}

func NewServiceCategoryRepository() interfaces_repositories.IServiceCategoryRepository {
	return ServiceCategoryRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}