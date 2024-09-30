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

type MessageHandler struct {
	UseCase usecase.MessageUseCase
}

func(handler *MessageHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateMessageRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertMessage(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *MessageHandler) FindAllMessage(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllMessage()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListMessageResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler MessageHandler) FindMessageById(ctx *gin.Context){
	 statusId, paramErr := uuid.Parse(ctx.Param("statusId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindMessageById(statusId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToMessageResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *MessageHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateMessageRequest

	statusId, paramErr := uuid.Parse(ctx.Param("statusId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateMessage(statusId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *MessageHandler) Remove(ctx *gin.Context){

	statusId, paramErr := uuid.Parse(ctx.Param("statusId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteMessage(statusId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}