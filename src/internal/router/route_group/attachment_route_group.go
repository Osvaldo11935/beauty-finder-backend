package route_group

import (
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAttachmentRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/attachment")

	r.POST("", handler.AttachmentHandler.Create)

	r.GET(":attachmentId", handler.AttachmentHandler.FindAttachmentById)
	r.GET("user/:userId", handler.AttachmentHandler.FindAttachmentByUserId)

	r.PUT(":attachmentId", handler.AttachmentHandler.Update)

	r.DELETE(":attachmentId", handler.AttachmentHandler.Remove)
}