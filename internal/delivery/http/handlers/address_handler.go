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

type AddressHandler struct {
	UseCase usecase.AddressUseCase
}

func(handler *AddressHandler) Create(ctx *gin.Context){

	userId := uuid.MustParse(ctx.Param("userId")) 
	
	var request models_requests_posts.CreateAddressRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertAddress(userId, request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func(handler *AddressHandler) CreateAddressAppointment(ctx *gin.Context){

	appointmentId := uuid.MustParse(ctx.Param("appointmentId")) 
	
	var request models_requests_posts.CreateAddressRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertAddressAppointment(appointmentId, request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func(handler AddressHandler) FindAddressByUserId(ctx *gin.Context){
	 userId, paramErr := uuid.Parse(ctx.Param("userId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindAddressByUserId(userId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToAddressResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}
func(handler AddressHandler) FindAddressByAppointmentId(ctx *gin.Context){
	appointmentId, paramErr := uuid.Parse(ctx.Param("appointmentId")) 

	if paramErr != nil {
	   ctx.JSON(http.StatusBadRequest, paramErr)
	   return
	}

	data, findErr := handler.UseCase.FindAddressByAppointmentId(appointmentId)

	if findErr != nil {
	   ctx.JSON(http.StatusBadRequest, findErr)
	   return
	}

	resp := models_responses.ToAddressResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler *AddressHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateAddressRequest

	userId, paramErr := uuid.Parse(ctx.Param("userId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateAddress(userId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *AddressHandler) Remove(ctx *gin.Context){

	userId, paramErr := uuid.Parse(ctx.Param("userId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteAddress(userId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}