package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type AppointmentRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAppointmentRepository(db *gorm.DB) interfaces_repositories.IAppointmentRepository{
	return AppointmentRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}