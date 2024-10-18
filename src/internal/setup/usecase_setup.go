package setup

import (
	"src/internal/usecase"
	"time"
)

type UseCaseSetup struct {
	UseCaseAddress usecase.AddressUseCase
	UseCaseAppointment usecase.AppointmentUseCase
	UseCaseAppointmentStatus usecase.AppointmentStatusUseCase
	UseCaseAttachment usecase.AttachmentUseCase
	UseCaseAttachmentType usecase.AttachmentTypeUseCase
	UseCaseServiceCategory usecase.ServiceCategoryUseCase
	UseCaseMessage usecase.MessageUseCase
	UseCasePerson usecase.PersonUseCase
	UseCaseRole usecase.RoleUseCase
	UseCaseService usecase.ServiceUseCase
	UseCaseServicePrice usecase.ServicePriceUseCase
	UseCaseUser usecase.UserUseCase
	UseCaseHttpClient usecase.HttpClientUseCase
	UseCaseFcmToken usecase.FcmTokenUseCase
	UseCaseRatingType usecase.RatingTypeUseCase
	UseCaseUserRating usecase.UserRatingUseCase
}

func NewUseCaseSetup(setup *RepositorySetup) *UseCaseSetup{
	return &UseCaseSetup{
		UseCaseAddress: usecase.AddressUseCase{Repo:  setup.AddressRepository},
		UseCaseAppointment: usecase.AppointmentUseCase{Repo: setup.AppointmentRepository},
		UseCaseAppointmentStatus: usecase.AppointmentStatusUseCase{Repo: setup.AppointmentStatusRepository},
		UseCaseAttachment: usecase.AttachmentUseCase{Repo: setup.AttachmentRepository},
		UseCaseAttachmentType: usecase.AttachmentTypeUseCase{Repo: setup.AttachmentTypeRepository},
	    UseCaseServiceCategory: usecase.ServiceCategoryUseCase{Repo: setup.ServiceCategoryRepository},
		UseCaseMessage: usecase.MessageUseCase{Repo: setup.MessageRepository},
		UseCasePerson: usecase.PersonUseCase{Repo: setup.PersonRepository, HttpClientUseCase: *usecase.NewHttpClientUseCase(3600 * time.Second)},
		UseCaseRole: usecase.RoleUseCase{Repo: setup.RoleRepository},
		UseCaseService: usecase.ServiceUseCase{Repo: setup.ServiceRepository},
		UseCaseServicePrice: usecase.ServicePriceUseCase{Repo: setup.ServicePriceRepository},
		UseCaseUser: usecase.UserUseCase{Repo: setup.PersonRepository},
		UseCaseUserRating: usecase.UserRatingUseCase{Repo: setup.UserRatingRepository},
		UseCaseRatingType: usecase.RatingTypeUseCase{Repo: setup.RatingTypeRepository},
		UseCaseHttpClient: *usecase.NewHttpClientUseCase(3600 * time.Second),
		UseCaseFcmToken: usecase.FcmTokenUseCase{ UserUseCase: usecase.UserUseCase{Repo: setup.PersonRepository}},
	}
}