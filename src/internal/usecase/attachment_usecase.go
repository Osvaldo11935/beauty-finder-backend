package usecase


import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type AttachmentUseCase struct {
	Repo interfaces_repositories.IAttachmentRepository
}


func(uc *AttachmentUseCase) InsertAttachment(request models_requests_posts.CreateAttachmentRequest) (*uuid.UUID, error){
	 req := entities.NewAttachment(request.Url, request.UserId)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateAttachmentError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *AttachmentUseCase) FindAttachmentByUserId(userId uuid.UUID) ([]entities.Attachment, error){
	
	var data []entities.Attachment

	findErr := uc.Repo.Query().
	        Preload("AttachmentType").
	        Find(&data, "UserId", userId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAttachmentError(findErr.Error())
	}

	return data, nil
}

func(uc *AttachmentUseCase) FindAttachmentById(AttachmentId uuid.UUID) (*entities.Attachment, error){
	var data entities.Attachment

	findErr := uc.Repo.Query().
			Preload("AttachmentType").
			First(&data, "ID", AttachmentId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAttachmentError(findErr.Error())
	}

	return &data, nil
}

func(uc *AttachmentUseCase) UpdateAttachment(AttachmentId uuid.UUID, request models_requests_puts.UpdateAttachmentRequest) (error){
	
	Attachment, findErr := uc.FindAttachmentById(AttachmentId)

	if findErr !=nil {
		return findErr
	}

	Attachment.Update(request.Url, request.UserId)
	
	updateErr := uc.Repo.Update(Attachment)

	if updateErr != nil {
	   return errors.UnknownUpdateAttachmentError(updateErr.Error())
	}

	return nil
}

func(uc *AttachmentUseCase) DeleteAttachment(AttachmentId uuid.UUID) error{
	
	Attachment, findErr := uc.FindAttachmentById(AttachmentId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(Attachment)

	if removeErr !=nil {
		return errors.UnknownDeleteAttachmentError(removeErr.Error())
	}

	return nil
}
