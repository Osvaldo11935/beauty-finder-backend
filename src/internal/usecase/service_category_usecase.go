package usecase

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type ServiceCategoryUseCase struct {
	Repo interfaces_repositories.IServiceCategoryRepository
}

func (uc *ServiceCategoryUseCase) InsertServiceCategory(request models_requests_posts.CreateServiceCategoryRequest) (*uuid.UUID, error) {
	req := entities.NewServiceCategory(request.Name, request.Description)

	createErr := uc.Repo.Insert(req)

	if createErr != nil {
		return nil, errors.UnknownCreateCategoryError(createErr.Error())
	}

	return &req.ID, nil
}

func (uc *ServiceCategoryUseCase) FindAllServiceCategory() ([]entities.ServiceCategory, error) {

	var data []entities.ServiceCategory

	findErr := uc.Repo.Query().
		Preload("Attachment").
		Find(&data).Error

	if findErr != nil {
		return nil, errors.UnknownFindCategoryError(findErr.Error())
	}

	return data, nil
}

func (uc *ServiceCategoryUseCase) FindServiceCategoryById(categoryId uuid.UUID) (*entities.ServiceCategory, error) {
	var data entities.ServiceCategory

	findErr := uc.Repo.Query().
		Preload("Attachment").
		First(&data, "ID", categoryId).Error

	if findErr != nil {
		return nil, errors.UnknownFindCategoryError(findErr.Error())
	}

	return &data, nil
}

func (uc *ServiceCategoryUseCase) UpdateServiceCategory(categoryId uuid.UUID, request models_requests_puts.UpdateServiceCategoryRequest) error {

	category, findErr := uc.FindServiceCategoryById(categoryId)

	if findErr != nil {
		return findErr
	}

	category.Update(request.Name, request.Description)

	updateErr := uc.Repo.Update(category)

	if updateErr != nil {
		return errors.UnknownUpdateCategoryError(updateErr.Error())
	}

	return nil
}

func (uc *ServiceCategoryUseCase) DeleteServiceCategory(categoryId uuid.UUID) error {

	category, findErr := uc.FindServiceCategoryById(categoryId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(category)

	if removeErr != nil {
		return errors.UnknownDeleteCategoryError(removeErr.Error())
	}

	return nil
}
