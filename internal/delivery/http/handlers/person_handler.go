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

type PersonHandler struct {
	UseCase usecase.PersonUseCase
}

func(handler *PersonHandler) Create(ctx *gin.Context){

	userId := uuid.MustParse(ctx.Param("userId")) 
	
	var request models_requests_posts.CreatePersonRequest

	 deserializerErr := ctx.ShouldBindJSON(&request)

	 if deserializerErr != nil{
        ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	 }

	 id, createErr := handler.UseCase.InsertPerson(userId, request)

	 if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return 
	 }

	 ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func(handler PersonHandler) FindPersonByUserId(ctx *gin.Context){
	 userId, paramErr := uuid.Parse(ctx.Param("userId")) 

	 if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	 }

	 data, findErr := handler.UseCase.FindPersonByUserId(userId)

	 if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	 }

	 resp := models_responses.ToPersonResponse(data)

	 ctx.JSON(http.StatusOK, resp)
}

func(handler PersonHandler) FindPersonByNationalRegistry(ctx *gin.Context){
	nationalRegistry := ctx.Param("nationalRegistry")

	data, findErr := handler.UseCase.FindPersonDataFromGovernmentApi(ctx, nationalRegistry)

	if findErr != nil {
	   ctx.JSON(http.StatusBadRequest, findErr)
	   return
	}
    result := models_responses.PersonDataFromGovernmentResponseToPersonResponse(&data.Data)

	ctx.JSON(http.StatusOK, result)
}

func(handler *PersonHandler) Update(ctx *gin.Context){

	var request models_requests_puts.UpdatePersonRequest

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

	updateErr := handler.UseCase.UpdatePerson(userId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func(handler *PersonHandler) Remove(ctx *gin.Context){

	userId, paramErr := uuid.Parse(ctx.Param("userId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeletePerson(userId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}