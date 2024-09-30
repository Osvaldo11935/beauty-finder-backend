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

type AttachmentTypeHandler struct {
	UseCase usecase.AttachmentTypeUseCase
}

func(handler *AttachmentTypeHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateAttachmentTypeRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertAttachmentType(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *AttachmentTypeHandler) FindAllAttachmentType(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllAttachmentType()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListAttachmentTypeResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler AttachmentTypeHandler) FindAttachmentTypeById(ctx *gin.Context){
	 statusId, paramErr := uuid.Parse(ctx.Param("statusId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindAttachmentTypeById(statusId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToAttachmentTypeResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *AttachmentTypeHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateAttachmentTypeRequest

	attachmentTypeId, paramErr := uuid.Parse(ctx.Param("attachmentTypeId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateAttachmentType(attachmentTypeId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *AttachmentTypeHandler) Remove(ctx *gin.Context){

	attachmentTypeId, paramErr := uuid.Parse(ctx.Param("attachmentTypeId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteAttachmentType(attachmentTypeId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}