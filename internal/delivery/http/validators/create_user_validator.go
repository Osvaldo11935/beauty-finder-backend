package validators

import (
	"fmt"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	"src/internal/domain/errors"
	"src/internal/persistence/database"
	"src/internal/persistence/repositories"
	"src/internal/usecase"
)

func ValidateCreateUser(user *models_requests_posts.CreateUserRequest) error {
	db, connectErr := database.Connect()

	if connectErr != nil {
		fmt.Println("Erro ao conectar no banco de dados no metodo de validação dos dados do usuario:", connectErr)
		return nil
	}

	userUseCase := usecase.UserUseCase{
		Repo: repositories.NewAddressRepository(db),
	}

	var validationErrors []string
	if user.Email != nil {

		validateEmailErr := errors.ValidateEmail(*user.Email)
		if validateEmailErr != nil {
			validationErrors = append(
				validationErrors,
				validateEmailErr.Error(),
			)
		}

		userByEmail, _, _ := userUseCase.FindUserByEmail(*user.Email)

		if userByEmail != nil {
			emailAlreadyExists := errors.EmailAlreadyExists()
			validationErrors = append(
				validationErrors,
				emailAlreadyExists.Error(),
			)
		}

	}

	validatePhoneNumberErr := errors.ValidatePhone(user.PhoneNumber)

	if validatePhoneNumberErr != nil {
		validationErrors = append(validationErrors, validatePhoneNumberErr.Error())
	}
	userByPhoneNumber, _, _ := userUseCase.FindUserByPhoneNumber(user.PhoneNumber)

	if userByPhoneNumber != nil {
		phoneNumberAlreadyExists := errors.PhoneNumberAlreadyExists()
		validationErrors = append(
			validationErrors,
			phoneNumberAlreadyExists.Error(),
		)
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return errors.ValidateCreateUserError(validationErrors)
}
