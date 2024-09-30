package usecase


import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type ServicePriceUseCase struct {
	Repo interfaces_repositories.IServicePriceRepository
}


func(uc *ServicePriceUseCase) InsertServicePrice(request models_requests_posts.CreateServicePriceRequest) (*uuid.UUID, error){
	 req := entities.NewServicePrice(request.Amount, request.ServiceId)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateServicePriceError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *ServicePriceUseCase) FindServicePriceByServiceId(serviceId uuid.UUID) ([]entities.ServicePrice, error){
	
	var data []entities.ServicePrice

	findErr := uc.Repo.Query().Find(&data, "ServiceId", serviceId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindServicePriceError(findErr.Error())
	}

	return data, nil
}

func(uc *ServicePriceUseCase) FindServicePriceById(servicePriceId uuid.UUID) (*entities.ServicePrice, error){
	var data entities.ServicePrice

	findErr := uc.Repo.Query().First(&data, "Id", servicePriceId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindServicePriceError(findErr.Error())
	}

	return &data, nil
}

func(uc *ServicePriceUseCase) UpdateServicePrice(servicePriceId uuid.UUID, request models_requests_puts.UpdateServicePriceRequest) (error){
	
	ServicePrice, findErr := uc.FindServicePriceById(servicePriceId)

	if findErr !=nil {
		return findErr
	}

	ServicePrice.Update(request.Amount, request.ServiceId)
	
	updateErr := uc.Repo.Update(ServicePrice)

	if updateErr != nil {
	   return errors.UnknownUpdateServicePriceError(updateErr.Error())
	}

	return nil
}

func(uc *ServicePriceUseCase) DeleteServicePrice(servicePriceId uuid.UUID) error{
	
	servicePrice, findErr := uc.FindServicePriceById(servicePriceId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(servicePrice)

	if removeErr !=nil {
		return errors.UnknownDeleteServicePriceError(removeErr.Error())
	}

	return nil
}
