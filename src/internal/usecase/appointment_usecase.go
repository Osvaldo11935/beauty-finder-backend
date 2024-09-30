package usecase


import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type AppointmentUseCase struct {
	Repo interfaces_repositories.IAppointmentRepository
}


func(uc *AppointmentUseCase) InsertAppointment(request models_requests_posts.CreateAppointmentRequest) (*uuid.UUID, error){
	 
	req := entities.NewAppointment(
		request.ProviderId, 
		request.ClientId, 
		request.ServiceId, 
		request.StartDate, 
		request.EndDate)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateAppointmentError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *AppointmentUseCase) FindAppointmentByProviderId(providerId uuid.UUID) ([]entities.Appointment, error){
	
	var data []entities.Appointment

	findErr := uc.Repo.Query().
	        Preload("Service").
			Preload("Status").
			Preload("Client").
	        Find(&data, "ProviderId", providerId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAppointmentError(findErr.Error())
	}

	return data, nil
}
func(uc *AppointmentUseCase) FindAppointmentByClientId(clientId uuid.UUID) ([]entities.Appointment, error){
	
	var data []entities.Appointment

	findErr := uc.Repo.Query().
	        Preload("Service").
			Preload("Status").
			Preload("Provider").
	        Find(&data, "ClientId", clientId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAppointmentError(findErr.Error())
	}

	return data, nil
}

func(uc *AppointmentUseCase) FindAppointmentById(Id uuid.UUID) (*entities.Appointment, error){
	var data entities.Appointment

	findErr := uc.Repo.Query().
			   Preload("Service").
			   Preload("Status").
			   Preload("Provider").
			   Preload("ClientId").
			   First(&data, "ID", Id).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAppointmentError(findErr.Error())
	}

	return &data, nil
}

func(uc *AppointmentUseCase) UpdateAppointment(Id uuid.UUID, request models_requests_puts.UpdateAppointmentRequest) (error){
	
	appointment, findErr := uc.FindAppointmentById(Id)

	if findErr !=nil {
		return findErr
	}

	appointment.Update(request.ServiceId)
	
	updateErr := uc.Repo.Update(appointment)

	if updateErr != nil {
	   return errors.UnknownUpdateAppointmentError(updateErr.Error())
	}

	return nil
}

func(uc *AppointmentUseCase) DeleteAppointment(Id uuid.UUID) error{
	
	appointment, findErr := uc.FindAppointmentById(Id)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(appointment)

	if removeErr !=nil {
		return errors.UnknownDeleteAppointmentError(removeErr.Error())
	}

	return nil
}
