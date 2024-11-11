package handlers

import (
	"net/http"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	"src/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DispatchNotificationHandler struct {
	FcmTokenUseCase usecase.FcmTokenUseCase
}


func (handler *DispatchNotificationHandler) DispatchStartConversationNotification(ctx *gin.Context) {
	
	userRequiredId := uuid.MustParse(ctx.Param("userRequiredId"))
	userApplicantId := uuid.MustParse(ctx.Param("userApplicantId"))
	

	handler.FcmTokenUseCase.DispatchStartConversationNotification(ctx, userApplicantId, userRequiredId)

	// if createErr != nil {
	//    ctx.JSON(http.StatusBadRequest, createErr)
	//    return
	// }

	ctx.JSON(http.StatusNoContent, nil)
}

func (handler *DispatchNotificationHandler) DispatchAssessmentNotification(ctx *gin.Context) {
	
	var request models_requests_posts.DispatchAssessmentNotification

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	userEvaluatorId := uuid.MustParse(ctx.Param("userEvaluatorId"))

	handler.FcmTokenUseCase.DispatchAssessmentNotification(ctx, userEvaluatorId, request)

	// if createErr != nil {
	//    ctx.JSON(http.StatusBadRequest, createErr)
	//    return
	// }

	ctx.JSON(http.StatusNoContent, nil)
}

func (handler *DispatchNotificationHandler) DispatchServiceNotification(ctx *gin.Context) {

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