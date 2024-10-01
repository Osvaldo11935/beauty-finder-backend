package usecase

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type UserUseCase struct {
	Repo interfaces_repositories.IUserRepository
}

func (uc *UserUseCase) InsertUser(roleId uuid.UUID, request models_requests_posts.CreateUserRequest) (*uuid.UUID, error) {
	req := entities.NewUser(
		request.Email,
		request.Email,
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
		First(&user, "Id", userId).Error

	if findErr != nil {
		return errors.UnknownCreateUserError(findErr.Error())
	}

	user.SetServicesProvided(serviceIds)
	
	uc.Repo.Remove(user.ServicesProvided)

	// if removeErr != nil {
	// 	return errors.UnknownDeleteUserError(removeErr.Error())
	// }

	createErr := uc.Repo.Insert(user.ServicesProvided)
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
		First(&data, "Id", UserId).Error

	if findErr != nil {
		return nil, errors.UnknownFindUserError(findErr.Error())
	}

	return &data, nil
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