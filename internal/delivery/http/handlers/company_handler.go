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

type CompanyHandler struct {
	UseCase usecase.CompanyUseCase
}

func (handler *CompanyHandler) Create(ctx *gin.Context) {
	var request models_requests_posts.CreateCompanyRequest

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	id, createErr := handler.UseCase.InsertCompany(request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}

func (handler *CompanyHandler) FindAllCompany(ctx *gin.Context) {

	data, findErr := handler.UseCase.FindAllCompany()

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToListCompanyResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func (handler CompanyHandler) FindCompanyById(ctx *gin.Context) {
	companyId, paramErr := uuid.Parse(ctx.Param("companyId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	data, findErr := handler.UseCase.FindCompanyByCompanyId(companyId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToCompanyResponse(data)

	ctx.JSON(http.StatusOK, resp)
}

func (handler CompanyHandler) FindCompanyByNationalRegistry(ctx *gin.Context) {
	nationalRegistry := ctx.Param("nationalRegistry")

	data, findErr := handler.UseCase.FindPersonDataByNationalRegistryFromGovernmentApi(ctx, nationalRegistry)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}
	result := models_responses.PersonDataFromGovernmentResponseToPersonResponse(&data.Data)

	ctx.JSON(http.StatusOK, result)
}

func (handler *CompanyHandler) Update(ctx *gin.Context) {

	var request models_requests_puts.UpdateCompanyRequest

	companyId, paramErr := uuid.Parse(ctx.Param("companyId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deserializeErr := ctx.ShouldBindJSON(request)

	if deserializeErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializeErr)
		return
	}

	updateErr := handler.UseCase.UpdateCompany(companyId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (handler *CompanyHandler) Remove(ctx *gin.Context) {

	companyId, paramErr := uuid.Parse(ctx.Param("companyId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteCompany(companyId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
