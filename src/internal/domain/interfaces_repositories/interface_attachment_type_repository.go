package interfaces_repositories

import (
	interfaces_repositories_common "src/internal/domain/interfaces_repositories/common"
	"gorm.io/gorm"
)

type IAttachmentTypeRepository interface {
	interfaces_repositories_common.IGormBaseRepository[gorm.DB]
}