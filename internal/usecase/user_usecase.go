package usecase

import (
	err "errors"
	"log"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"
	"src/internal/domain/object_values"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NearbyUser struct {
	User     entities.User
	Distance float64
}

type UserUseCase struct {
	Repo interfaces_repositories.IUserRepository
}

func (uc *UserUseCase) InsertUser(roleId uuid.UUID, request models_requests_posts.CreateUserRequest) (*uuid.UUID, error) {
	req := entities.NewUser(
		request.Email,
		request.UserName,
		request.Password,
		request.PhoneNumber,
		roleId)

	createErr := uc.Repo.Insert(req)

	if createErr != nil {
		return nil, errors.UnknownCreateUserError(createErr.Error())
	}

	return &req.ID, nil
}

func (uc *UserUseCase) InsertServiceProvided(userId uuid.UUID, serviceIds []uuid.UUID) error {

	var user entities.User

	findErr := uc.Repo.Query().
		Preload("ServicesProvided").
		Where("Id", userId).
		Where("RoleId", object_values.ROLE_SERVICE_PROVIDER_ID).
		First(&user).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return errors.NotFoundServiceProviderError()
		}
		return errors.UnknownCreateUserError(findErr.Error())
	}

	user.SetServicesProvided(serviceIds)

	uc.Repo.Remove(user.ServicesProvided)

	createErr := uc.Repo.Insert(user.ServicesProvided)
	if createErr != nil {
		return errors.UnknownDeleteUserError(createErr.Error())
	}
	return nil
}

func (uc *UserUseCase) InsertFcmToken(userId uuid.UUID, fcmToken models_requests_posts.CreateFcmTokenRequest) error {

	var user entities.User

	findErr := uc.Repo.Query().
		Preload("FcmToken").
		Where("Id", userId).
		First(&user).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return errors.NotFoundUserError()
		}
		return errors.UnknownCreateUserError(findErr.Error())
	}

	user.SetFcmToken(
		fcmToken.FcmToken,
		fcmToken.DeviceName,
		fcmToken.DeviceId,
	)

	uc.Repo.Remove(user.FcmToken)

	createErr := uc.Repo.Insert(user.FcmToken)
	if createErr != nil {
		return errors.UnknownDeleteUserError(createErr.Error())
	}
	return nil
}

func (uc *UserUseCase) FindAllUser() ([]entities.User, error) {

	var data []entities.User

	findErr := uc.Repo.Query().
		Preload("Role").
		Preload("Person").
		Find(&data).Error

	if findErr != nil {
		return nil, errors.UnknownFindUserError(findErr.Error())
	}

	return data, nil
}

func (uc *UserUseCase) FindUserByServiceId(serviceId uuid.UUID) ([]entities.User, error) {

	var users []entities.User

	var data []entities.ServiceProvider

	findErr := uc.Repo.Query().
		Preload("Service").
		Preload("Provider.Person").
		Preload("Provider.Role").
		Preload("Provider.FcmToken").
		Find(&data, "ServiceId", serviceId).Error

	if findErr != nil {
		return nil, errors.UnknownFindUserError(findErr.Error())
	}

	for _, d := range data {
		users = append(users, *d.Provider)
	}

	return users, nil
}

func (uc *UserUseCase) FindUserById(UserId uuid.UUID) (*entities.User, error) {
	var data entities.User

	findErr := uc.Repo.Query().
		Preload("Role").
		Preload("Person").
		Preload("FcmToken").
		First(&data, "Id", UserId).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundUserError()
		}
		return nil, errors.UnknownFindUserError(findErr.Error())
	}

	return &data, nil
}

func (uc *UserUseCase) FindUserByCredentials(email string, password string) (*entities.User, *[]string, error) {

	var claims []string

	var userData entities.User

	findUserErr := uc.Repo.Query().
		Preload("Role").
		Where("Email", email).
		Where("Password", password).First(&userData).Error

	if findUserErr != nil {
		log.Printf("Erro ao buscar dados do usuario: %v", findUserErr)
		return nil, nil, errors.InvalidCredentialError()
	}

	roleData := userData.Role
	claims = append(claims, roleData.Name)

	return &userData, &claims, nil
}

func (uc *UserUseCase) FindUserByPhoneNumber(phoneNumber string) (*entities.User, *[]string, error) {

	var claims []string

	var userData entities.User

	findUserErr := uc.Repo.Query().
		Preload("Role").
		Where("PhoneNumber", phoneNumber).
		First(&userData).Error

	if findUserErr != nil {
		if err.Is(findUserErr, gorm.ErrRecordNotFound) {
			return nil, nil, errors.NotFoundUserError()
		}
		log.Printf("Erro ao buscar dados do usuario: %v", findUserErr)
		return nil, nil, errors.InvalidCredentialError()
	}

	roleData := userData.Role
	claims = append(claims, roleData.Name)

	return &userData, &claims, nil
}

func (uc *UserUseCase) FindUserByEmail(email string) (*entities.User, *[]string, error) {

	var claims []string

	var userData entities.User

	findUserErr := uc.Repo.Query().
		Preload("Role").
		Where("Email", email).
		First(&userData).Error

	if findUserErr != nil {
		if err.Is(findUserErr, gorm.ErrRecordNotFound) {
			return nil, nil, errors.NotFoundUserError()
		}
		log.Printf("Erro ao buscar dados do usuario: %v", findUserErr)
		return nil, nil, errors.InvalidCredentialError()
	}

	roleData := userData.Role
	claims = append(claims, roleData.Name)

	return &userData, &claims, nil
}

func (uc *UserUseCase) FindUsersNearBy(serviceId uuid.UUID, latitude float64, longitude float64, radiusKm float64) ([]*NearbyUser, error) {

	var serviceProviders []entities.ServiceProvider

	if err := uc.Repo.Query().
		Preload("Provider.Address").
		Preload("Provider.Role").
		Preload("Provider.Person").
		Preload("Provider.FcmToken").
		Find(&serviceProviders).Error; err != nil {
		return nil, err
	}

	var nearbyUsers []*NearbyUser

	for _, serviceProvider := range serviceProviders {

		distance := Haversine(latitude, longitude, serviceProvider.Provider.Address.Latitude, serviceProvider.Provider.Address.Longitude)

		if distance <= radiusKm {

			nearbyUser := NearbyUser{
				User:     *serviceProvider.Provider,
				Distance: distance,
			}

			nearbyUsers = append(nearbyUsers, &nearbyUser)
		}
	}

	return nearbyUsers, nil
}

func (uc *UserUseCase) UpdateUser(UserId uuid.UUID, request models_requests_puts.UpdateUserRequest) error {

	User, findErr := uc.FindUserById(UserId)

	if findErr != nil {
		return findErr
	}

	User.Update(
		request.Email,
		request.Email,
		request.Password,
		request.PhoneNumber,
		request.RoleId)

	updateErr := uc.Repo.Update(User)

	if updateErr != nil {
		return errors.UnknownUpdateUserError(updateErr.Error())
	}

	return nil
}

func (uc *UserUseCase) DeleteUser(UserId uuid.UUID) error {

	User, findErr := uc.FindUserById(UserId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(User)

	if removeErr != nil {
		return errors.UnknownDeleteUserError(removeErr.Error())
	}

	return nil
}
