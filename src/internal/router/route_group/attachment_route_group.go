package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAttachmentRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/attachment")

	r.POST("", middlewares.Auth(), handler.AttachmentHandler.Create)

	r.GET(":attachmentId", middlewares.Auth(), handler.AttachmentHandler.FindAttachmentById)
	r.GET("user/:userId", middlewares.Auth(), handler.AttachmentHandler.FindAttachmentByUserId)

	r.PUT(":attachmentId", middlewares.Auth(), handler.AttachmentHandler.Update)

	r.DELETE(":attachmentId", middlewares.Auth(), handler.AttachmentHandler.Remove)
}