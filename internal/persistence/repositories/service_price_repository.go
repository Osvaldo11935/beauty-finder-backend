package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type ServicePriceRepository struct {
	*repositories_common.GormBaseRepository
}

func NewServicePriceRepository(db *gorm.DB) interfaces_repositories.IServicePriceRepository {
	return ServicePriceRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}