package validators

import (
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	"src/internal/domain/errors"
)

func ValidateCreateUser(user *models_requests_posts.CreateUserRequest) error {
	var validationErrors []string

	validateEmailErr := errors.ValidateEmail(*user.Email)
	validatePhoneNumberErr := errors.ValidatePhone(user.PhoneNumber)

	if validateEmailErr != nil {
		validationErrors = append(
			validationErrors,
			validateEmailErr.Error(),
		)
	}

	if validatePhoneNumberErr != nil {
		validationErrors = append(validationErrors, validatePhoneNumberErr.Error())
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return errors.ValidateCreateUserError(validationErrors)
}
