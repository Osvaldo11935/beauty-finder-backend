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

type ServiceHandler struct {
	UseCase usecase.ServiceUseCase
}

func(handler *ServiceHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateServiceRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertService(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *ServiceHandler) FindAllService(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllService()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListServiceResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler *ServiceHandler) FindServiceByProviderId(ctx *gin.Context){
      
	providerId := uuid.MustParse(ctx.Param("providerId"))

	data, findErr := handler.UseCase.FindServiceByProviderId(providerId)
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListServiceResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler *ServiceHandler) FindServiceByCategoryId(ctx *gin.Context){
     
	categoryId := uuid.MustParse(ctx.Param("categoryId"))

	data, findErr := handler.UseCase.FindServiceByCategoryId(categoryId)
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListServiceResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler ServiceHandler) FindServiceById(ctx *gin.Context){
	 serviceId, paramErr := uuid.Parse(ctx.Param("serviceId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindServiceById(serviceId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToServiceResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *ServiceHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateServiceRequest

	serviceId, paramErr := uuid.Parse(ctx.Param("serviceId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateService(serviceId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *ServiceHandler) Remove(ctx *gin.Context){

	serviceId, paramErr := uuid.Parse(ctx.Param("serviceId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteService(serviceId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}