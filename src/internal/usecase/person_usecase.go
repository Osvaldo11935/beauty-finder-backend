package usecase

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/interfaces_repositories"
	"src/internal/domain/errors"
	"github.com/google/uuid"
)

type PersonUseCase struct {
	Repo interfaces_repositories.IPersonRepository
}

func (uc *PersonUseCase) InsertPerson(userId uuid.UUID, request models_requests_posts.CreatePersonRequest) (*uuid.UUID, error) {
	req := entities.NewPerson(
			request.FullName, 
			request.BirthDate, 
			request.Gender, 
			request.Naturalness, 
			request.MaritalStatus, 
			request.FatherName, 
			request.MotherName, 
			request.NationalRegistry, 
			request.PlaceIssuanceDocument, 
			request.DateIssueDocument, 
			&userId)

	createErr := uc.Repo.Insert(&req)

	if createErr != nil {
		return nil, errors.UnknownCreatePersonError(createErr.Error())
	}

	return &req.ID, nil
}

func (uc *PersonUseCase) FindPersonByUserId(userId uuid.UUID) (*entities.Person, error) {
	var data entities.Person

	findErr := uc.Repo.Query().
		First(&data, "UserId", userId).Error

	if findErr != nil {
		return nil, errors.UnknownFindPersonError(findErr.Error())
	}

	return &data, nil
}

func (uc *PersonUseCase) UpdatePerson(userId uuid.UUID, request models_requests_puts.UpdatePersonRequest) error {

	Person, findErr := uc.FindPersonByUserId(userId)

	if findErr != nil {
		return findErr
	}

	Person.Update(
		request.FullName, 
		request.BirthDate, 
		request.Gender, 
		request.Naturalness, 
		request.MaritalStatus, 
		request.FatherName, 
		request.MotherName, 
		request.NationalRegistry, 
		request.PlaceIssuanceDocument, 
		request.DateIssueDocument, 
		request.UserId)

	updateErr := uc.Repo.Update(Person)

	if updateErr != nil {
		return errors.UnknownUpdatePersonError(updateErr.Error())
	}

	return nil
}

func (uc *PersonUseCase) DeletePerson(userId uuid.UUID) error {

	person, findErr := uc.FindPersonByUserId(userId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(person)

	if removeErr != nil {
		return errors.UnknownDeletePersonError(removeErr.Error())
	}

	return nil
}
