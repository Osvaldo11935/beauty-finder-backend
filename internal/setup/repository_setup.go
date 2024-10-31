package setup

import (
	"src/internal/domain/interfaces_repositories"
	"src/internal/persistence/repositories"

	"gorm.io/gorm"
)

type RepositorySetup struct {
	AddressRepository           interfaces_repositories.IAddressRepository
	AppointmentRepository       interfaces_repositories.IAppointmentRepository
	AppointmentStatusRepository interfaces_repositories.IAppointmentStatusRepository
	AttachmentRepository        interfaces_repositories.IAttachmentRepository
	AttachmentTypeRepository    interfaces_repositories.IAttachmentRepository
	ServiceCategoryRepository   interfaces_repositories.IServiceCategoryRepository
	MessageRepository           interfaces_repositories.IMessageRepository
	PersonRepository            interfaces_repositories.IPersonRepository
	RoleRepository              interfaces_repositories.IRoleRepository
	ServiceRepository           interfaces_repositories.IServiceRepository
	ServicePriceRepository      interfaces_repositories.IServicePriceRepository
	UserRepository              interfaces_repositories.IUserRepository
	RatingTypeRepository        interfaces_repositories.IRatingTypeRepository
	UserRatingRepository        interfaces_repositories.IUserRatingRepository
}

func NewRepositorySetup(db *gorm.DB) *RepositorySetup {
	return &RepositorySetup{
		AddressRepository:           repositories.NewAddressRepository(db),
		AppointmentRepository:       repositories.NewAppointmentRepository(db),
		AppointmentStatusRepository: repositories.NewAppointmentStatusRepository(db),
		AttachmentRepository:        repositories.NewAttachmentRepository(db),
		AttachmentTypeRepository:    repositories.NewAttachmentTypeRepository(db),
		ServiceCategoryRepository:   repositories.NewServiceCategoryRepository(db),
		MessageRepository:           repositories.NewMessageRepository(db),
		PersonRepository:            repositories.NewPersonRepository(db),
		RoleRepository:              repositories.NewRoleRepository(db),
		ServiceRepository:           repositories.NewServiceRepository(db),
		ServicePriceRepository:      repositories.NewServicePriceRepository(db),
		UserRepository:              repositories.NewUserRepository(db),
		RatingTypeRepository: repositories.NewRatingTypeRepository(db),
		UserRatingRepository: repositories.NewUserRatingRepository(db),
	}
}
