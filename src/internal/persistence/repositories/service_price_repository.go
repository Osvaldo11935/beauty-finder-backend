package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type ServicePriceRepository struct {
	*repositories_common.GormBaseRepository
}

func NewServicePriceRepository() interfaces_repositories.IServicePriceRepository {
	return ServicePriceRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}