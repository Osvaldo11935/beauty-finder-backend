package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type AddressRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAddressRepository(db *gorm.DB) interfaces_repositories.IAddressRepository {
	return AddressRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}