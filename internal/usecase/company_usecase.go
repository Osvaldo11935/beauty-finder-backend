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

	err "errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyUseCase struct {
	HttpClientUseCase HttpClientUseCase
	Repo              interfaces_repositories.ICompanyRepository
}

func (uc *CompanyUseCase) InsertCompany(request models_requests_posts.CreateCompanyRequest) (*uuid.UUID, error) {
	req := entities.NewCompany(request)

	createErr := uc.Repo.Insert(&req)

	if createErr != nil {
		return nil, errors.UnknownCreateCompanyError(createErr.Error())
	}

	return &req.ID, nil
}

func (uc *CompanyUseCase) FindAllCompany() ([]*entities.Company, error) {
	var data []*entities.Company

	findErr := uc.Repo.Query().
		Find(&data).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundFindCompanyError()
		}
		return nil, errors.UnknownFindCompanyError(findErr.Error())
	}

	return data, nil
}

func (uc *CompanyUseCase) FindCompanyByCompanyId(companyId uuid.UUID) (*entities.Company, error) {
	var data entities.Company

	findErr := uc.Repo.Query().
		First(&data, "Id", companyId).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundFindCompanyError()
		}
		return nil, errors.UnknownFindCompanyError(findErr.Error())
	}

	return &data, nil
}

func (uc *CompanyUseCase) FindPersonDataByNationalRegistryFromGovernmentApi(ctx context.Context, nationalRegistry string) (*models_responses.PersonByNationalRegistry, error) {
	var apiResponse models_responses.PersonByNationalRegistry

	config, configErr := configs.LoadConfig()

	if configErr != nil {
		fmt.Println("Erro ao carregar as configurações:", configErr)
		return nil, configErr
	}

	headers := map[string]string{"Content-Type": "application/json"}

	response, requestExternalApiErr := uc.HttpClientUseCase.Get(ctx, config.QueryingPersonDataCompany+"="+nationalRegistry, headers)

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

func (uc *CompanyUseCase) UpdateCompany(companyId uuid.UUID, request models_requests_puts.UpdateCompanyRequest) error {

	company, findErr := uc.FindCompanyByCompanyId(companyId)

	if findErr != nil {
		return findErr
	}

	company.Update(request)

	updateErr := uc.Repo.Update(company)

	if updateErr != nil {
		return errors.UnknownUpdateCompanyError(updateErr.Error())
	}

	return nil
}

func (uc *CompanyUseCase) DeleteCompany(companyId uuid.UUID) error {

	company, findErr := uc.FindCompanyByCompanyId(companyId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(company)

	if removeErr != nil {
		return errors.UnknownDeleteCompanyError(removeErr.Error())
	}

	return nil
}
