package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewMessageRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/message")

	r.POST("", handler.MessageHandler.Create)

	r.GET("", handler.MessageHandler.FindAllMessage)
	r.GET(":messageId", handler.MessageHandler.FindMessageById)

	r.PUT(":messageId", handler.MessageHandler.Update)

	r.DELETE(":messageId", handler.MessageHandler.Remove)
}