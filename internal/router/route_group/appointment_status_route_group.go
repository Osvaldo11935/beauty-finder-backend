package route_group

import (
	"src/internal/delivery/http/handlers"
	"src/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func NewAppointmentStatusRouteGroup(r *gin.RouterGroup, handler handlers.AppointmentStatusHandler){
     route := r.Group("/appointment-status")

	 route.POST("", middlewares.Auth(), handler.Create)

	 route.GET("", middlewares.Auth(), handler.FindAllAppointmentStatus)
	 route.GET(":statusId", middlewares.Auth(), handler.FindAppointmentStatusById)

	 route.PUT(":statusId", middlewares.Auth(), handler.Update)
	 
	 route.DELETE(":statusId", middlewares.Auth(), handler.Remove)     
}