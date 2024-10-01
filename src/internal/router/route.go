package router

import (
	"src/internal/router/route_group"
	"src/internal/setup"
	"src/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRoute(setup *setup.HandlerSetup, pool *usecase.Pool) *gin.Engine {
	router := gin.Default()

	baseRoute := router.Group("/api")

	route_group.NewAppointmentRouteGroup(baseRoute, *setup)
	route_group.NewAppointmentStatusRouteGroup(baseRoute, setup.AppointmentStatusHandler)
	route_group.NewAttachmentRouteGroup(baseRoute, *setup)
	route_group.NewAttachmentTypeRouteGroup(baseRoute, *setup)
	route_group.NewCategoryRouteGroup(baseRoute, *setup)
	route_group.NewMessageRouteGroup(baseRoute, *setup)
	route_group.NewPersonRouteGroup(baseRoute, *setup)
	route_group.NewRoleRouteGroup(baseRoute, *setup)
	route_group.NewServicePriceRouteGroup(baseRoute, *setup)
	route_group.NewServiceRouteGroup(baseRoute, *setup)
	route_group.NewUserRouteGroup(baseRoute, *setup)
	route_group.NewChatRouteGroup(baseRoute, &setup.UserHandler.UseCase, pool)

	return router
}