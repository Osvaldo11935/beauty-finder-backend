package route_group

import (
	"src/internal/delivery/websocket/handlers"
	"src/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewWebSocketRouteGroup(route *gin.RouterGroup, setup *usecase.AppointmentUseCase) {
	r := route.Group("/ws")

	r.GET("appointment", func(ctx *gin.Context) {
		handlers.FindAppointmentWebsocketHandler(ctx, *setup)
	})
}
