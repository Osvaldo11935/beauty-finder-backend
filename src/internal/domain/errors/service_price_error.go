package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateServicePriceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar preço do serviço", description, nil)
}
func UnknownFindServicePriceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar preço do serviço", description, nil)
}
func UnknownDeleteServicePriceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar preço do serviço", description, nil)
}
func UnknownUpdateServicePriceError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar preço do serviço", description, nil)
}
func ValidateCreateServicePriceError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar preço do serviço", "", err)
}