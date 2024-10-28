package usecase


import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"
	"github.com/google/uuid"
)

type RatingTypeUseCase struct {
	Repo interfaces_repositories.IRatingTypeRepository
}


func(uc *RatingTypeUseCase) InsertRatingType(request models_requests_posts.CreateRatingTypeRequest) (*uuid.UUID, error){
	 req := entities.NewRatingType(request)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateRatingTypeError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *RatingTypeUseCase) FindAllRatingType() ([]entities.RatingType, error){
	
	var data []entities.RatingType

	findErr := uc.Repo.Query().Find(&data).Error

	if findErr != nil {
	   return nil, errors.UnknownFindRatingTypeError(findErr.Error())
	}

	return data, nil
}

func(uc *RatingTypeUseCase) FindRatingTypeById(ratingTypeId uuid.UUID) (*entities.RatingType, error){
	var data entities.RatingType

	findErr := uc.Repo.Query().First(&data, "ID", ratingTypeId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindRatingTypeError(findErr.Error())
	}

	return &data, nil
}

func(uc *RatingTypeUseCase) UpdateRatingType(ratingTypeId uuid.UUID, request models_requests_puts.UpdateRatingTypeRequest) (error){
	
	RatingType, findErr := uc.FindRatingTypeById(ratingTypeId)

	if findErr !=nil {
		return findErr
	}

	RatingType.Update(request)
	
	updateErr := uc.Repo.Update(RatingType)

	if updateErr != nil {
	   return errors.UnknownUpdateRatingTypeError(updateErr.Error())
	}

	return nil
}

func(uc *RatingTypeUseCase) DeleteRatingType(ratingTypeId uuid.UUID) error{
	
	RatingType, findErr := uc.FindRatingTypeById(ratingTypeId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(RatingType)

	if removeErr !=nil {
		return errors.UnknownDeleteRatingTypeError(removeErr.Error())
	}

	return nil
}
