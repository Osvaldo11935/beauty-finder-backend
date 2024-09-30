package usecase


import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type AttachmentTypeUseCase struct {
	Repo interfaces_repositories.IAttachmentTypeRepository
}


func(uc *AttachmentTypeUseCase) InsertAttachmentType(request models_requests_posts.CreateAttachmentTypeRequest) (*uuid.UUID, error){
	 req := entities.NewAttachmentType(request.Type, request.Description)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateAttachmentTypeError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *AttachmentTypeUseCase) FindAllAttachmentType() ([]entities.AttachmentType, error){
	
	var data []entities.AttachmentType

	findErr := uc.Repo.Query().Find(&data).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAttachmentTypeError(findErr.Error())
	}

	return data, nil
}

func(uc *AttachmentTypeUseCase) FindAttachmentTypeById(AttachmentTypeId uuid.UUID) (*entities.AttachmentType, error){
	var data entities.AttachmentType

	findErr := uc.Repo.Query().First(&data, "ID", AttachmentTypeId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAttachmentTypeError(findErr.Error())
	}

	return &data, nil
}

func(uc *AttachmentTypeUseCase) UpdateAttachmentType(AttachmentTypeId uuid.UUID, request models_requests_puts.UpdateAttachmentTypeRequest) (error){
	
	AttachmentType, findErr := uc.FindAttachmentTypeById(AttachmentTypeId)

	if findErr !=nil {
		return findErr
	}

	AttachmentType.Update(request.Type, request.Description)
	
	updateErr := uc.Repo.Update(AttachmentType)

	if updateErr != nil {
	   return errors.UnknownUpdateAttachmentTypeError(updateErr.Error())
	}

	return nil
}

func(uc *AttachmentTypeUseCase) DeleteAttachmentType(AttachmentTypeId uuid.UUID) error{
	
	AttachmentType, findErr := uc.FindAttachmentTypeById(AttachmentTypeId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(AttachmentType)

	if removeErr !=nil {
		return errors.UnknownDeleteAttachmentTypeError(removeErr.Error())
	}

	return nil
}
