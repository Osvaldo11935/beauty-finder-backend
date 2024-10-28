package usecase


import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type ServiceUseCase struct {
	Repo interfaces_repositories.IServiceRepository
}


func(uc *ServiceUseCase) InsertService(request models_requests_posts.CreateServiceRequest) (*uuid.UUID, error){
	 req := entities.NewService(request.Name, request.Description, request.CategoryId)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateServiceError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *ServiceUseCase) FindAllService() ([]entities.Service, error){
	
	var data []entities.Service

	findErr := uc.Repo.Query().
	 		Preload("Category").
			Preload("Price").
			Preload("Attachment").
			Find(&data).Error

	if findErr != nil {
	   return nil, errors.UnknownFindServiceError(findErr.Error())
	}

	return data, nil
}
func(uc *ServiceUseCase) FindServiceByProviderId(providerId uuid.UUID) ([]entities.Service, error){
	
	var services []entities.Service

	var data []entities.ServiceProvider

	findErr := uc.Repo.Query().
	 		Preload("Service.Category").
			Preload("Service.Attachment").
			Preload("Service.Price").
			Find(&data, "ProviderId", providerId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindServiceError(findErr.Error())
	}
    
	for _, d := range(data){
		services = append(services, *d.Service)
	}

	return services, nil
}
func(uc *ServiceUseCase) FindServiceByCategoryId(categoryId uuid.UUID) ([]entities.Service, error){
	
	var data []entities.Service

	findErr := uc.Repo.Query().
	 		Preload("Category").
			Preload("Price").
			Preload("Attachment").
			Find(&data, "CategoryId", categoryId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindServiceError(findErr.Error())
	}

	return data, nil
}

func(uc *ServiceUseCase) FindServiceById(ServiceId uuid.UUID) (*entities.Service, error){
	var data entities.Service

	findErr := uc.Repo.Query().
	        Preload("Category").
			Preload("Price").
			Preload("Attachment").
			First(&data, "Id", ServiceId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindServiceError(findErr.Error())
	}

	return &data, nil
}

func(uc *ServiceUseCase) UpdateService(ServiceId uuid.UUID, request models_requests_puts.UpdateServiceRequest) (error){
	
	Service, findErr := uc.FindServiceById(ServiceId)

	if findErr !=nil {
		return findErr
	}

	Service.Update(request.Name, request.Description, request.CategoryId)
	
	updateErr := uc.Repo.Update(Service)

	if updateErr != nil {
	   return errors.UnknownUpdateServiceError(updateErr.Error())
	}

	return nil
}

func(uc *ServiceUseCase) DeleteService(ServiceId uuid.UUID) error{
	
	Service, findErr := uc.FindServiceById(ServiceId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(Service)

	if removeErr !=nil {
		return errors.UnknownDeleteServiceError(removeErr.Error())
	}

	return nil
}
