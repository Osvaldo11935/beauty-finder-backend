package errors

import (
	"regexp"
	error_common "src/internal/domain/errors/common"

	"github.com/asaskevich/govalidator"
)

func ValidateEmail(email string) error {
	if !govalidator.IsEmail(email) {
		return error_common.NewCustomError(error_common.ERR_INVALID_EMAIL, "E-mail inválido", "O endereço de e-mail fornecido não é válido.", nil)
	}
	return nil
}

func EmailAlreadyExists() error {
	return error_common.NewCustomError(error_common.ERR_EXISTING_EMAIL, "Falha ao cadastrar usuario", "E-mail informado ja existe.", nil)
}
func PhoneNumberAlreadyExists() error {
	return error_common.NewCustomError(error_common.ERR_EXISTING_PHONENUMBER, "Falha ao cadastrar usuario", "Numero de telefone informado ja existe.", nil)
}
func ValidatePhone(phoneNumber string) error {
	re := regexp.MustCompile(`^\d{3}\d{3}\d{3}$`)
	if !re.MatchString(phoneNumber) {
		return error_common.NewCustomError(error_common.ERR_INVALID_PHONE_NUMBER, "Número de telefone inválido", "O número de telefone fornecido não está no formato 999 999 999.", nil)
	}
	return nil
}
