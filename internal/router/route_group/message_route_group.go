package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewMessageRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/message")

	r.POST("", middlewares.Auth(), handler.MessageHandler.Create)

	r.GET("", middlewares.Auth(), handler.MessageHandler.FindAllMessage)
	r.GET(":messageId", middlewares.Auth(), handler.MessageHandler.FindMessageById)

	r.PUT(":messageId", middlewares.Auth(), handler.MessageHandler.Update)

	r.DELETE(":messageId", middlewares.Auth(), handler.MessageHandler.Remove)
}