package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type RoleRepository struct {
	*repositories_common.GormBaseRepository
}

func NewRoleRepository(db *gorm.DB) interfaces_repositories.IRoleRepository {
	return RoleRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}