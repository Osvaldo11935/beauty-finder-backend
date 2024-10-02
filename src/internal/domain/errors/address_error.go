package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateAddressError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar endereço", description, nil)
}
func UnknownFindAddressError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar endereço", description, nil)
}
func UnknownDeleteAddressError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar endereço", description, nil)
}
func UnknownUpdateAddressError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar endereço", description, nil)
}
func ValidateCreateAddressError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar endereço", "", err)
}