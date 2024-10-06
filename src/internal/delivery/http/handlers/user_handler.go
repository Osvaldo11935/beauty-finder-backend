package handlers

import (
	"net/http"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	models_responses "src/internal/delivery/http/models/responses"
	"src/internal/domain/entities"
	"src/internal/domain/object_values"
	"src/internal/security"
	"src/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UseCase usecase.UserUseCase
}

func (handler *UserHandler) CreateAdmin(ctx *gin.Context) {
	var request models_requests_posts.CreateUserRequest

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	id, createErr := handler.UseCase.InsertUser(object_values.ROLE_ADMIN_ID, request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *UserHandler) CreateClient(ctx *gin.Context) {
	var request models_requests_posts.CreateUserRequest

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	id, createErr := handler.UseCase.InsertUser(object_values.ROLE_CLIENT_ID, request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *UserHandler) CreateServiceProvider(ctx *gin.Context) {

	body, paramExists := ctx.Get("request")

	if !paramExists {
		ctx.JSON(http.StatusBadRequest, "Request data not found")
		return
	}

	request, serializerErr := body.(models_requests_posts.CreateUserRequest)

	if !serializerErr {
		ctx.JSON(http.StatusBadRequest, "Failed to parse request")
		return
	}

	id, createErr := handler.UseCase.InsertUser(object_values.ROLE_SERVICE_PROVIDER_ID, request)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusOK, models_responses.NewCreateResponse(*id))
}
func (handler *UserHandler) CreateServiceProvided(ctx *gin.Context) {

	userId := uuid.MustParse(ctx.Param("userId"))

	var request models_requests_posts.CreateServiceProvidedRequest

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	createErr := handler.UseCase.InsertServiceProvided(userId, request.ServiceIds)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
func (handler *UserHandler) FindAllUser(ctx *gin.Context) {

	data, findErr := handler.UseCase.FindAllUser()

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToListUserResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler UserHandler) FindUserByServiceId(ctx *gin.Context) {
	serviceId, paramErr := uuid.Parse(ctx.Param("serviceId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	data, findErr := handler.UseCase.FindUserByServiceId(serviceId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToListUserResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler UserHandler) FindUsersNearBy(ctx *gin.Context) {

	serviceId, _ := uuid.Parse(ctx.Param("serviceId"))
	lat, _ := strconv.ParseFloat(ctx.Param("latitude"), 64)
	long, _ := strconv.ParseFloat(ctx.Param("longitude"), 64)
	radiusKm, _ := strconv.ParseFloat(ctx.Query("radiusKm"), 64)

	data, findErr := handler.UseCase.FindUsersNearBy(serviceId, lat, long, radiusKm)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
func (handler UserHandler) FindUserById(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	data, findErr := handler.UseCase.FindUserById(uuid.MustParse(userId.(string)))

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	resp := models_responses.ToUserResponse(data)

	ctx.JSON(http.StatusOK, resp)
}
func (handler UserHandler) FindToken(ctx *gin.Context) {
	var findErr error
	var claims *[]string
	var user *entities.User

	jwtSecurity := security.NewJwtTokenService()
	var request models_requests_posts.GetTokenRequest

	deserializerErr := ctx.ShouldBindJSON(&request)

	if deserializerErr != nil {
		ctx.JSON(http.StatusBadRequest, deserializerErr)
		return
	}

	if request.PhoneNumber != nil {
		user, claims, findErr = handler.UseCase.FindUserByPhoneNumber(*request.PhoneNumber)
	} else {
		user, claims, findErr = handler.UseCase.FindUserByCredentials(*request.Email, *request.Password)
	}

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	token, generateTokenErr := jwtSecurity.GenerateToken(*user.Email, user.ID.String(), *claims)

	if generateTokenErr != nil {
		ctx.JSON(http.StatusBadRequest, generateTokenErr)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
func (handler *UserHandler) Update(ctx *gin.Context) {

	var request models_requests_puts.UpdateUserRequest

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

	updateErr := handler.UseCase.UpdateUser(userId, request)

	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, updateErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
func (handler *UserHandler) Remove(ctx *gin.Context) {

	userId, paramErr := uuid.Parse(ctx.Param("userId"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, paramErr)
		return
	}

	deleteErr := handler.UseCase.DeleteUser(userId)

	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, deleteErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
