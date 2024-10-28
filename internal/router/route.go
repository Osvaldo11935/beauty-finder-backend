package router

import (
	"src/internal/router/route_group"
	"src/internal/setup"
	"src/internal/usecase"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRoute(setup *setup.HandlerSetup, pool *usecase.Pool) *gin.Engine {
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	baseRoute := router.Group("/api")
    route_group.NewAddressRouteGroup(baseRoute, *setup)
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
    route_group.NewWebSocketRouteGroup(baseRoute, &setup.AppointmentHandler.UseCase);
	route_group.NewUserRatingRouteGroup(baseRoute, *setup)
	route_group.NewRatingTypeRouteGroup(baseRoute, *setup)
	return router
}
