package setup

import (
	"src/internal/delivery/http/handlers"
)

type HandlerSetup struct {
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
}

func NewHandlerSetup(setup *UseCaseSetup) *HandlerSetup{
	 return &HandlerSetup{
		AppointmentHandler: handlers.AppointmentHandler{UseCase: setup.UseCaseAppointment},
		AppointmentStatusHandler: handlers.AppointmentStatusHandler{UseCase: setup.UseCaseAppointmentStatus},
		AttachmentHandler:        handlers.AttachmentHandler{UseCase: setup.UseCaseAttachment},
		AttachmentTypeHandler:    handlers.AttachmentTypeHandler{UseCase: setup.UseCaseAttachmentType},
		ServiceCategoryHandler:   handlers.ServiceCategoryHandler{UseCase: setup.UseCaseServiceCategory},
		MessageHandler:           handlers.MessageHandler{UseCase: setup.UseCaseMessage},
		PersonHandler:            handlers.PersonHandler{UseCase: setup.UseCasePerson},
		RoleHandler:              handlers.RoleHandler{UseCase: setup.UseCaseRole},
		ServiceHandler:           handlers.ServiceHandler{UseCase: setup.UseCaseService},
		ServicePriceHandler:      handlers.ServicePriceHandler{UseCase: setup.UseCaseServicePrice},
		UserHandler:              handlers.UserHandler{UseCase: setup.UseCaseUser},
	 }
}