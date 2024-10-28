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

type ServiceCategoryHandler struct {
	UseCase usecase.ServiceCategoryUseCase
}

func(handler *ServiceCategoryHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateServiceCategoryRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertServiceCategory(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *ServiceCategoryHandler) FindAllServiceCategory(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllServiceCategory()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListServiceCategoryResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler ServiceCategoryHandler) FindServiceCategoryById(ctx *gin.Context){
	 statusId, paramErr := uuid.Parse(ctx.Param("statusId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindServiceCategoryById(statusId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToServiceCategoryResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *ServiceCategoryHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateServiceCategoryRequest

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

	updateErr := handler.UseCase.UpdateServiceCategory(statusId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *ServiceCategoryHandler) Remove(ctx *gin.Context){

	statusId, paramErr := uuid.Parse(ctx.Param("statusId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteServiceCategory(statusId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}