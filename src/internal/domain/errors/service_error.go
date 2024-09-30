package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateServiceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar serviço", description, nil)
}
func UnknownFindServiceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar serviço", description, nil)
}
func UnknownDeleteServiceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar serviço", description, nil)
}
func UnknownUpdateServiceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar serviço", description, nil)
}
func ValidateCreateServiceError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar serviço", "", err)
}