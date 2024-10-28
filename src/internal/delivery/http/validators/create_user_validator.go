package validators

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	"src/internal/domain/errors"
	"src/internal/persistence/repositories"
	"src/internal/usecase"
)

func ValidateCreateUser(user *models_requests_posts.CreateUserRequest) error {
	userUseCase := usecase.UserUseCase{
		Repo: repositories.NewAddressRepository(),
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
