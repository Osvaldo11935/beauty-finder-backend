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

type ServicePriceHandler struct {
	UseCase usecase.ServicePriceUseCase
}

func(handler *ServicePriceHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateServicePriceRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertServicePrice(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *ServicePriceHandler) FindServicePriceByServiceId(ctx *gin.Context){
      
	serviceId := uuid.MustParse(ctx.Param("serviceId"))

	data, findErr := handler.UseCase.FindServicePriceByServiceId(serviceId)
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListServicePriceResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler ServicePriceHandler) FindServicePriceById(ctx *gin.Context){
	priceId, paramErr := uuid.Parse(ctx.Param("priceId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindServicePriceById(priceId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToServicePriceResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *ServicePriceHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateServicePriceRequest

	priceId, paramErr := uuid.Parse(ctx.Param("priceId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateServicePrice(priceId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *ServicePriceHandler) Remove(ctx *gin.Context){

	priceId, paramErr := uuid.Parse(ctx.Param("priceId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteServicePrice(priceId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}