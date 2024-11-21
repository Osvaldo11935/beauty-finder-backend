package usecase

import (
	err "errors"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleUseCase struct {
	Repo interfaces_repositories.IRoleRepository
}

func (uc *RoleUseCase) InsertRole(request models_requests_posts.CreateRoleRequest) (*uuid.UUID, error) {
	req := entities.NewRole(request.Name)

	createErr := uc.Repo.Insert(req)

	if createErr != nil {
		return nil, errors.UnknownCreateRoleError(createErr.Error())
	}

	return &req.ID, nil
}

func (uc *RoleUseCase) FindAllRole() ([]entities.Role, error) {

	var data []entities.Role

	findErr := uc.Repo.Query().Find(&data).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundFindRoleError()
		}
		return nil, errors.UnknownFindRoleError(findErr.Error())
	}

	return data, nil
}

func (uc *RoleUseCase) FindRoleById(roleId uuid.UUID) (*entities.Role, error) {
	var data entities.Role

	findErr := uc.Repo.Query().First(&data, "ID", roleId).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundFindRoleError()
		}
		return nil, errors.UnknownFindRoleError(findErr.Error())
	}

	return &data, nil
}

func (uc *RoleUseCase) UpdateRole(roleId uuid.UUID, request models_requests_puts.UpdateRoleRequest) error {

	role, findErr := uc.FindRoleById(roleId)

	if findErr != nil {
		return findErr
	}

	role.Update(request.Name)

	updateErr := uc.Repo.Update(role)

	if updateErr != nil {
		return errors.UnknownUpdateRoleError(updateErr.Error())
	}

	return nil
}

func (uc *RoleUseCase) DeleteRole(roleId uuid.UUID) error {

	role, findErr := uc.FindRoleById(roleId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(role)

	if removeErr != nil {
		return errors.UnknownDeleteRoleError(removeErr.Error())
	}

	return nil
}
