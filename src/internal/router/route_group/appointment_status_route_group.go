package route_group

import (
	"src/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func NewAppointmentStatusRouteGroup(r *gin.RouterGroup, handler handlers.AppointmentStatusHandler){
     route := r.Group("/appointment-status")

	 route.POST("", handler.Create)

	 route.GET("", handler.FindAllAppointmentStatus)
	 route.GET(":statusId", handler.FindAppointmentStatusById)

	 route.PUT(":statusId", handler.Update)
	 
	 route.DELETE(":statusId", handler.Remove)     
}