package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type PersonRepository struct {
	*repositories_common.GormBaseRepository
}

func NewPersonRepository(db *gorm.DB) interfaces_repositories.IPersonRepository {
	return PersonRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}