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

type AppointmentHandler struct {
	UseCase         usecase.AppointmentUseCase
	FcmTokenUseCase usecase.FcmTokenUseCase
}

func (handler *AppointmentHandler) Create(ctx *gin.Context) {
	var request models_requests_posts.CreateAppointmentRequest

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	id, createErr := handler.UseCase.InsertAppointment(request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *AppointmentHandler) DispatchServiceNotification(ctx *gin.Context) {

	var request models_requests_posts.DispatchServiceNotification

	serviceId := uuid.MustParse(ctx.Param("serviceId"))
	appointmentId := uuid.MustParse(ctx.Param("appointmentId"))

	paramErr := ctx.ShouldBindJSON(&request)

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	handler.FcmTokenUseCase.DispatchServiceNotification(ctx, serviceId, appointmentId, request)

	ctx.JSON(http.StatusNoContent, nil)
}
func (handler *AppointmentHandler) FindAppointmentByClientId(ctx *gin.Context) {

	clientId := uuid.MustParse(ctx.Param("clientId"))

	data, findErr := handler.UseCase.FindAppointmentByClientId(clientId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToListAppointmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func (handler *AppointmentHandler) FindAppointmentByProviderId(ctx *gin.Context) {

	providerId := uuid.MustParse(ctx.Param("providerId"))

	data, findErr := handler.UseCase.FindAppointmentByProviderId(providerId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToListAppointmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func (handler AppointmentHandler) FindAppointmentById(ctx *gin.Context) {
	appointmentId, paramErr := uuid.Parse(ctx.Param("appointmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	data, findErr := handler.UseCase.FindAppointmentById(appointmentId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToAppointmentResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func (handler *AppointmentHandler) Update(ctx *gin.Context) {

	var request models_requests_puts.UpdateAppointmentRequest

	appointmentId, paramErr := uuid.Parse(ctx.Param("appointmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateAppointment(appointmentId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
func (handler *AppointmentHandler) SetProviderAppointment(ctx *gin.Context) {

	var request models_requests_puts.UpdateAppointmentRequest

	appointmentId, paramErr := uuid.Parse(ctx.Param("appointmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(&request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.SetProviderAppointment(appointmentId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
func (handler *AppointmentHandler) Remove(ctx *gin.Context) {

	appointmentId, paramErr := uuid.Parse(ctx.Param("appointmentId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteAppointment(appointmentId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
