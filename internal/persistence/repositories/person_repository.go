package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type PersonRepository struct {
	*repositories_common.GormBaseRepository
}

func NewPersonRepository() interfaces_repositories.IPersonRepository {
	return PersonRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}