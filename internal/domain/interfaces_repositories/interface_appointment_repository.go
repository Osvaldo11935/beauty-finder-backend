package interfaces_repositories

import (
	"gorm.io/gorm"
	interfaces_repositories_common "src/internal/domain/interfaces_repositories/common"
)

type IAppointmentRepository interface {
	interfaces_repositories_common.IGormBaseRepository[gorm.DB]
}