package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"src/internal/configs"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	models_responses "src/internal/delivery/http/models/responses"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type PersonUseCase struct {
	HttpClientUseCase HttpClientUseCase
	Repo              interfaces_repositories.IPersonRepository
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

func (uc *PersonUseCase) FindPersonDataFromGovernmentApi(ctx context.Context, nationalRegistry string) (*models_responses.PersonByNationalRegistry, error) {
	var apiResponse models_responses.PersonByNationalRegistry

	config, configErr := configs.LoadConfig()
	
	if configErr != nil {
		fmt.Println("Erro ao carregar as configurações:", configErr)
		return nil, configErr
	}

	headers := map[string]string{"Content-Type": "application/json"}

	response, requestExternalApiErr := uc.HttpClientUseCase.Get(ctx, config.QueryingPersonData+"="+nationalRegistry, headers)

	if requestExternalApiErr != nil {
		return nil, errors.UnknownFindPersonError(requestExternalApiErr.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, readBodyErr := ioutil.ReadAll(response.Body)
	if readBodyErr != nil {
		return nil, errors.UnknownFindPersonError(readBodyErr.Error())

	}

	if deserializeErr := json.Unmarshal(body, &apiResponse); deserializeErr != nil {
		return nil, errors.UnknownFindPersonError(deserializeErr.Error())
	}
	return &apiResponse, nil
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
