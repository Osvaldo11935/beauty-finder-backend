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

type AppointmentStatusHandler struct {
	UseCase usecase.AppointmentStatusUseCase
}

func(handler *AppointmentStatusHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateAppointmentStatusRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertAppointmentStatus(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *AppointmentStatusHandler) FindAllAppointmentStatus(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllAppointmentStatus()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListAppointmentStatusResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler AppointmentStatusHandler) FindAppointmentStatusById(ctx *gin.Context){
	 statusId, paramErr := uuid.Parse(ctx.Param("statusId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindAppointmentStatusById(statusId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToAppointmentStatusResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *AppointmentStatusHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateAppointmentStatusRequest

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

	updateErr := handler.UseCase.UpdateAppointmentStatus(statusId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *AppointmentStatusHandler) Remove(ctx *gin.Context){

	statusId, paramErr := uuid.Parse(ctx.Param("statusId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteAppointmentStatus(statusId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}