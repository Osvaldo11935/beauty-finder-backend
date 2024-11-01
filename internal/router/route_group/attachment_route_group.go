package route_group

import (
	"src/internal/delivery/http/middlewares"
	"src/internal/setup"

	"github.com/gin-gonic/gin"
)

func NewAttachmentRouteGroup(route *gin.RouterGroup, handler setup.HandlerSetup) {
	r := route.Group("/attachment")

	r.POST("user/:userId/attachmentType/:attachmentTypeId", handler.AttachmentHandler.CreateAttachmentUser)
	r.POST("service/:serviceId/attachmentType/:attachmentTypeId", middlewares.Auth(), handler.AttachmentHandler.CreateAttachmentService)
	r.POST("category/:categoryId/attachmentType/:attachmentTypeId", middlewares.Auth(), handler.AttachmentHandler.CreateAttachmentCategory)

	r.GET(":attachmentId", middlewares.Auth(), handler.AttachmentHandler.FindAttachmentById)
	r.GET("user/:userId", middlewares.Auth(), handler.AttachmentHandler.FindAttachmentByUserId)
	r.GET("category/:categoryId", middlewares.Auth(), handler.AttachmentHandler.FindAttachmentByCategoryId)
	r.GET("service/:serviceId", middlewares.Auth(), handler.AttachmentHandler.FindAttachmentByServiceId)
	r.GET("file/:fileId", handler.AttachmentHandler.Download)

	r.PUT(":attachmentId", middlewares.Auth(), handler.AttachmentHandler.Update)

	r.DELETE(":attachmentId", middlewares.Auth(), handler.AttachmentHandler.Remove)
}
