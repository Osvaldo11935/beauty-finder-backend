package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"

	"gorm.io/gorm"
)

type AppointmentStatusRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAppointmentStatusRepository(db *gorm.DB) interfaces_repositories.IAppointmentStatusRepository{
	return AppointmentStatusRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository(db).(*repositories_common.GormBaseRepository),
	}
}