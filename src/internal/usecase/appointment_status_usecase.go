package usecase

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type AppointmentStatusUseCase struct {
	Repo interfaces_repositories.IAppointmentStatusRepository
}


func(uc *AppointmentStatusUseCase) InsertAppointmentStatus(request models_requests_posts.CreateAppointmentStatusRequest) (*uuid.UUID, error){
	 req := entities.NewAppointmentStatus(request.Type, request.Description)

	 createErr := uc.Repo.Insert(&req)

	 if createErr != nil {
		return nil, errors.UnknownCreateStatusError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *AppointmentStatusUseCase) FindAllAppointmentStatus() ([]entities.AppointmentStatus, error){
	
	var data []entities.AppointmentStatus

	findErr := uc.Repo.Query().Find(&data).Error

	if findErr != nil {
	   return nil, errors.UnknownFindStatusError(findErr.Error())
	}

	return data, nil
}

func(uc *AppointmentStatusUseCase) FindAppointmentStatusById(statusId uuid.UUID) (*entities.AppointmentStatus, error){
	var data entities.AppointmentStatus

	findErr := uc.Repo.Query().First(&data, "ID", statusId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindStatusError(findErr.Error())
	}

	return &data, nil
}

func(uc *AppointmentStatusUseCase) UpdateAppointmentStatus(statusId uuid.UUID, request models_requests_puts.UpdateAppointmentStatusRequest) (error){
	
	status, findErr := uc.FindAppointmentStatusById(statusId)

	if findErr !=nil {
		return findErr
	}

	status.Update(request.Type, request.Description)
	
	updateErr := uc.Repo.Update(status)

	if updateErr != nil {
	   return errors.UnknownUpdateStatusError(updateErr.Error())
	}

	return nil
}

func(uc *AppointmentStatusUseCase) DeleteAppointmentStatus(statusId uuid.UUID) error{
	
	status, findErr := uc.FindAppointmentStatusById(statusId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(status)

	if removeErr !=nil {
		return errors.UnknownDeleteStatusError(removeErr.Error())
	}

	return nil
}
