package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type AddressRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAddressRepository() interfaces_repositories.IAddressRepository {
	return AddressRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}