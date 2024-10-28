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

type RatingTypeHandler struct {
	UseCase usecase.RatingTypeUseCase
}

func(handler *RatingTypeHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateRatingTypeRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertRatingType(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *RatingTypeHandler) FindAllRatingType(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllRatingType()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListRatingTypeResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler RatingTypeHandler) FindRatingTypeById(ctx *gin.Context){
	 ratingTypeId, paramErr := uuid.Parse(ctx.Param("ratingTypeId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindRatingTypeById(ratingTypeId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToRatingTypeResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *RatingTypeHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateRatingTypeRequest

	ratingTypeId, paramErr := uuid.Parse(ctx.Param("ratingTypeId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateRatingType(ratingTypeId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *RatingTypeHandler) Remove(ctx *gin.Context){

	ratingTypeId, paramErr := uuid.Parse(ctx.Param("ratingTypeId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteRatingType(ratingTypeId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}