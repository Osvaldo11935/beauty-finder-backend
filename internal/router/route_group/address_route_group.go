package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAddressRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/address")

	r.POST("/user/:userId", handler.AddressHandler.Create)
	r.POST("/appointment/:appointmentId", handler.AddressHandler.CreateAddressAppointment)

	r.GET("/user/:userId", middlewares.Auth(), handler.AddressHandler.FindAddressByUserId)
	r.GET("/appointment/:appointmentId", middlewares.Auth(), handler.AddressHandler.FindAddressByAppointmentId)

	r.PUT("/user/:userId", middlewares.Auth(), handler.AddressHandler.Update)

	r.DELETE("/user/:userId", middlewares.Auth(), handler.AddressHandler.Remove)
}
