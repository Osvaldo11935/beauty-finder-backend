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

type UserRatingHandler struct {
	UseCase usecase.UserRatingUseCase
}

func(handler *UserRatingHandler) Create(ctx *gin.Context){
	 var request models_requests_posts.CreateUserRatingRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertUserRating(request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *UserRatingHandler) FindAllUserRating(ctx *gin.Context){
      
	data, findErr := handler.UseCase.FindAllUserRating()
	
	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
    
	resp := models_responses.ToListUserRatingResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func(handler UserRatingHandler) FindUserRatingById(ctx *gin.Context){
	 userId, paramErr := uuid.Parse(ctx.Param("userId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindUserRatingById(userId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToUserRatingResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler *UserRatingHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdateUserRatingRequest

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

	updateErr := handler.UseCase.UpdateUserRating(userId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *UserRatingHandler) Remove(ctx *gin.Context){

	userId, paramErr := uuid.Parse(ctx.Param("userId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteUserRating(userId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}