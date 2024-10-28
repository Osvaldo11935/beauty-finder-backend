package setup

import (
	"src/internal/delivery/http/handlers"
)

type HandlerSetup struct {
	AddressHandler           handlers.AddressHandler
	AppointmentHandler       handlers.AppointmentHandler
	AppointmentStatusHandler handlers.AppointmentStatusHandler
	AttachmentHandler        handlers.AttachmentHandler
	AttachmentTypeHandler    handlers.AttachmentTypeHandler
	ServiceCategoryHandler   handlers.ServiceCategoryHandler
	MessageHandler           handlers.MessageHandler
	PersonHandler            handlers.PersonHandler
	RoleHandler              handlers.RoleHandler
	ServiceHandler           handlers.ServiceHandler
	ServicePriceHandler      handlers.ServicePriceHandler
	UserHandler              handlers.UserHandler
	RatingTypeHandler        handlers.RatingTypeHandler
	UserRatingHandler        handlers.UserRatingHandler
}

func NewHandlerSetup(setup *UseCaseSetup, otherSetup *OtherSetup) *HandlerSetup {
	return &HandlerSetup{
		AddressHandler: handlers.AddressHandler{UseCase: setup.UseCaseAddress},
		AppointmentHandler: handlers.AppointmentHandler{UseCase: setup.UseCaseAppointment,
			FcmTokenUseCase: setup.UseCaseFcmToken,
		},
		AppointmentStatusHandler: handlers.AppointmentStatusHandler{UseCase: setup.UseCaseAppointmentStatus},
		AttachmentHandler: handlers.AttachmentHandler{UseCase: setup.UseCaseAttachment,
			FileManagerService: otherSetup.FileManager,
		},
		AttachmentTypeHandler:  handlers.AttachmentTypeHandler{UseCase: setup.UseCaseAttachmentType},
		ServiceCategoryHandler: handlers.ServiceCategoryHandler{UseCase: setup.UseCaseServiceCategory},
		MessageHandler:         handlers.MessageHandler{UseCase: setup.UseCaseMessage},
		PersonHandler:          handlers.PersonHandler{UseCase: setup.UseCasePerson},
		RoleHandler:            handlers.RoleHandler{UseCase: setup.UseCaseRole},
		ServiceHandler:         handlers.ServiceHandler{UseCase: setup.UseCaseService},
		ServicePriceHandler:    handlers.ServicePriceHandler{UseCase: setup.UseCaseServicePrice},
		UserHandler:            handlers.UserHandler{UseCase: setup.UseCaseUser},
		RatingTypeHandler:      handlers.RatingTypeHandler{UseCase: setup.UseCaseRatingType},
		UserRatingHandler:      handlers.UserRatingHandler{UseCase: setup.UseCaseUserRating, FcmTokenUseCase: setup.UseCaseFcmToken},
	}
}
