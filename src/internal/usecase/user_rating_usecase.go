package usecase

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type UserRatingUseCase struct {
	Repo interfaces_repositories.IUserRatingRepository
}

func (uc *UserRatingUseCase) InsertUserRating(request models_requests_posts.CreateUserRatingRequest) (*uuid.UUID, error) {
	req := entities.NewUserRating(request)

	createErr := uc.Repo.Insert(req)

	if createErr != nil {
		return nil, errors.UnknownCreateUserRatingError(createErr.Error())
	}

	return &req.ID, nil
}

func (uc *UserRatingUseCase) FindAllUserRating() ([]entities.UserRating, error) {

	var data []entities.UserRating

	findErr := uc.Repo.Query().
		Preload("UserEvaluator").
		Preload("UserAvaluated").
		Find(&data).Error

	if findErr != nil {
		return nil, errors.UnknownFindUserRatingError(findErr.Error())
	}

	return data, nil
}

func (uc *UserRatingUseCase) FindUserRatingById(UserRatingId uuid.UUID) (*entities.UserRating, error) {
	var data entities.UserRating

	findErr := uc.Repo.Query().
		Preload("UserEvaluator").
		Preload("UserAvaluated").
		First(&data, "ID", UserRatingId).Error

	if findErr != nil {
		return nil, errors.UnknownFindUserRatingError(findErr.Error())
	}

	return &data, nil
}

func (uc *UserRatingUseCase) UpdateUserRating(UserRatingId uuid.UUID, request models_requests_puts.UpdateUserRatingRequest) error {

	UserRating, findErr := uc.FindUserRatingById(UserRatingId)

	if findErr != nil {
		return findErr
	}

	UserRating.Update(request)

	updateErr := uc.Repo.Update(UserRating)

	if updateErr != nil {
		return errors.UnknownUpdateUserRatingError(updateErr.Error())
	}

	return nil
}

func (uc *UserRatingUseCase) DeleteUserRating(UserRatingId uuid.UUID) error {

	UserRating, findErr := uc.FindUserRatingById(UserRatingId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(UserRating)

	if removeErr != nil {
		return errors.UnknownDeleteUserRatingError(removeErr.Error())
	}

	return nil
}
