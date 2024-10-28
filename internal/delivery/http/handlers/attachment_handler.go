package handlers

import (
	"net/http"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	models_responses "src/internal/delivery/http/models/responses"
	service_interface "src/internal/services/interface_services"
	"src/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttachmentHandler struct {
	UseCase            usecase.AttachmentUseCase
	FileManagerService service_interface.IFileManager
}

func (handler *AttachmentHandler) CreateAttachmentUser(ctx *gin.Context) {

	userId := uuid.MustParse(ctx.Param("userId"))

	attachmentTypeId := uuid.MustParse(ctx.Param("attachmentTypeId"))

	url, uploadFileErr := handler.FileManagerService.Upload(ctx)

	if uploadFileErr != nil {
		ctx.JSON(http.StatusBadRequest, uploadFileErr)
		return
	}

	request := models_requests_posts.CreateAttachmentRequest{
		Url:              *url,
		UserId:           &userId,
		AttachmentTypeId: attachmentTypeId,
	}

	id, createErr := handler.UseCase.InsertAttachment(request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *AttachmentHandler) CreateAttachmentService(ctx *gin.Context) {

	serviceId := uuid.MustParse(ctx.Param("serviceId"))

	attachmentTypeId := uuid.MustParse(ctx.Param("attachmentTypeId"))

	url, uploadFileErr := handler.FileManagerService.Upload(ctx)

	if uploadFileErr != nil {
		ctx.JSON(http.StatusBadRequest, uploadFileErr)
		return
	}

	request := models_requests_posts.CreateAttachmentRequest{
		Url:              *url,
		ServiceId:        &serviceId,
		AttachmentTypeId: attachmentTypeId,
	}

	id, createErr := handler.UseCase.InsertAttachment(request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *AttachmentHandler) CreateAttachmentCategory(ctx *gin.Context) {

	categoryId := uuid.MustParse(ctx.Param("categoryId"))

	attachmentTypeId := uuid.MustParse(ctx.Param("attachmentTypeId"))

	url, uploadFileErr := handler.FileManagerService.Upload(ctx)

	if uploadFileErr != nil {
		ctx.JSON(http.StatusBadRequest, uploadFileErr)
		return
	}

	request := models_requests_posts.CreateAttachmentRequest{
		Url:              *url,
		CategoryId:       &categoryId,
		AttachmentTypeId: attachmentTypeId,
	}

	id, createErr := handler.UseCase.InsertAttachment(request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *AttachmentHandler) FindAttachmentByUserId(ctx *gin.Context) {

	userId := uuid.MustParse(ctx.Param("userId"))

	data, findErr := handler.UseCase.FindAttachmentByUserId(userId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToListAttachmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler *AttachmentHandler) FindAttachmentByCategoryId(ctx *gin.Context) {

	categoryId := uuid.MustParse(ctx.Param("categoryId"))

	data, findErr := handler.UseCase.FindAttachmentByCategoryId(categoryId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToAttachmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler *AttachmentHandler) FindAttachmentByServiceId(ctx *gin.Context) {

	serviceId := uuid.MustParse(ctx.Param("serviceId"))

	data, findErr := handler.UseCase.FindAttachmentByServiceId(serviceId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToAttachmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler AttachmentHandler) FindAttachmentById(ctx *gin.Context) {
	attachmentId, paramErr := uuid.Parse(ctx.Param("attachmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	data, findErr := handler.UseCase.FindAttachmentById(attachmentId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToAttachmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler AttachmentHandler) Download(ctx *gin.Context) {
	fileId := ctx.Param("fileId")

	data, findErr := handler.FileManagerService.Download(fileId)

	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data_uri": data,
	})
}
func (handler *AttachmentHandler) Update(ctx *gin.Context) {

	var request models_requests_puts.UpdateAttachmentRequest

	attachmentId, paramErr := uuid.Parse(ctx.Param("attachmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateAttachment(attachmentId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
func (handler *AttachmentHandler) Remove(ctx *gin.Context) {

	attachmentId, paramErr := uuid.Parse(ctx.Param("attachmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteAttachment(attachmentId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
