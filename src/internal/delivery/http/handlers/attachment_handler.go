package handlers

import (
	"net/http"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	models_responses "src/internal/delivery/http/models/responses"
	"src/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttachmentHandler struct {
	UseCase usecase.AttachmentUseCase
}

func(handler *AttachmentHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateAttachmentRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertAttachment(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *AttachmentHandler) FindAttachmentByUserId(ctx *gin.Context){
      
	userId := uuid.MustParse(ctx.Param("userId"))

	data, findErr := handler.UseCase.FindAttachmentByUserId(userId)
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListAttachmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler AttachmentHandler) FindAttachmentById(ctx *gin.Context){
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

func(handler *AttachmentHandler) Update(ctx *gin.Context){

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

func(handler *AttachmentHandler) Remove(ctx *gin.Context){

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