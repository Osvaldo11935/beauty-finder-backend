package repositories

import (
	"src/internal/domain/interfaces_repositories"
	repositories_common "src/internal/persistence/repositories/common"
)

type AppointmentRepository struct {
	*repositories_common.GormBaseRepository
}

func NewAppointmentRepository() interfaces_repositories.IAppointmentRepository{
	return AppointmentRepository{
		GormBaseRepository: repositories_common.NewGormBaseRepository().(*repositories_common.GormBaseRepository),
	}
}