package setup

import (
	"src/internal/domain/interfaces_repositories"
	"src/internal/persistence/repositories"
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
}

func NewRepositorySetup() *RepositorySetup {
	return &RepositorySetup{
		AddressRepository:           repositories.NewAddressRepository(),
		AppointmentRepository:       repositories.NewAppointmentRepository(),
		AppointmentStatusRepository: repositories.NewAppointmentStatusRepository(),
		AttachmentRepository:        repositories.NewAttachmentRepository(),
		AttachmentTypeRepository:    repositories.NewAttachmentTypeRepository(),
		ServiceCategoryRepository:   repositories.NewServiceCategoryRepository(),
		MessageRepository:           repositories.NewMessageRepository(),
		PersonRepository:            repositories.NewPersonRepository(),
		RoleRepository:              repositories.NewRoleRepository(),
		ServiceRepository:           repositories.NewServiceRepository(),
		ServicePriceRepository:      repositories.NewServicePriceRepository(),
		UserRepository:              repositories.NewUserRepository(),
	}
}
