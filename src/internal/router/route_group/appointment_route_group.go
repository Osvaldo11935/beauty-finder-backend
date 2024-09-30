package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAppointmentRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
     r := route.Group("/appointment")

	 r.POST("", handler.AppointmentHandler.Create)

	 r.GET(":appointmentId", handler.AppointmentHandler.FindAppointmentById)
	 r.GET("client/:clientId", handler.AppointmentHandler.FindAppointmentByClientId)
	 r.GET("provider/:providerId", handler.AppointmentHandler.FindAppointmentByProviderId)

	 r.PUT(":appointmentId", handler.AppointmentHandler.Update)
	 
	 r.DELETE(":appointmentId", handler.AppointmentHandler.Remove)
}