package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAppointmentRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/appointment")

	r.POST("", middlewares.Auth(), handler.AppointmentHandler.Create)

	r.GET(":appointmentId", middlewares.Auth(), handler.AppointmentHandler.FindAppointmentById)
	r.GET("client/:clientId", middlewares.Auth(), handler.AppointmentHandler.FindAppointmentByClientId)
	r.GET("provider/:providerId", middlewares.Auth(), handler.AppointmentHandler.FindAppointmentByProviderId)

	r.PUT(":appointmentId", middlewares.Auth(), handler.AppointmentHandler.Update)
	
	r.PATCH(":appointmentId/provider", middlewares.Auth(), handler.AppointmentHandler.SetProviderAppointment)
	r.PATCH(":appointmentId/status-type/:statusTypeId", middlewares.Auth(), handler.AppointmentHandler.UpdateStatusTypeAppointment)

	r.DELETE(":appointmentId", middlewares.Auth(), handler.AppointmentHandler.Remove)
}
