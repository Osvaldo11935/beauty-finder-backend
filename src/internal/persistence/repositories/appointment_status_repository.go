package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type AppointmentStatusRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAppointmentStatusRepository() interfaces_repositories.IAppointmentStatusRepository{
	return AppointmentStatusRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}