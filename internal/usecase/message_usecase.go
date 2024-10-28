package usecase

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type MessageUseCase struct {
	Repo interfaces_repositories.IMessageRepository
}


func(uc *MessageUseCase) InsertMessage(request models_requests_posts.CreateMessageRequest) (*uuid.UUID, error){
	 req := entities.NewMessage(
			request.Type, 
			request.Body,
			request.SenderId,
			request.ReceiverId,
	   )

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateStatusError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *MessageUseCase) FindAllMessage() ([]entities.Message, error){
	
	var data []entities.Message

	findErr := uc.Repo.Query().Find(&data).Error

	if findErr != nil {
	   return nil, errors.UnknownFindStatusError(findErr.Error())
	}

	return data, nil
}

func(uc *MessageUseCase) FindMessageById(messageId uuid.UUID) (*entities.Message, error){
	var data entities.Message

	findErr := uc.Repo.Query().First(&data, "ID", messageId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindStatusError(findErr.Error())
	}

	return &data, nil
}

func(uc *MessageUseCase) UpdateMessage(messageId uuid.UUID, request models_requests_puts.UpdateMessageRequest) (error){
	
	status, findErr := uc.FindMessageById(messageId)

	if findErr !=nil {
		return findErr
	}

	status.Update(request.Type, request.Body, request.SenderId, request.ReceiverId)
	
	updateErr := uc.Repo.Update(status)

	if updateErr != nil {
	   return errors.UnknownUpdateStatusError(updateErr.Error())
	}

	return nil
}

func(uc *MessageUseCase) DeleteMessage(messageId uuid.UUID) error{
	
	status, findErr := uc.FindMessageById(messageId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(status)

	if removeErr !=nil {
		return errors.UnknownDeleteStatusError(removeErr.Error())
	}

	return nil
}
