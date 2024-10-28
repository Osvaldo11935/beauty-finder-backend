package errors

import (
	error_common "src/internal/domain/errors/common"
)

func UnknownCreateStatusError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar estado", description, nil)
}
func UnknownFindStatusError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar estado", description, nil)
}
func UnknownDeleteStatusError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar estado", description, nil)
}
func UnknownUpdateStatusError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar estado", description, nil)
}
func ValidateCreateStatusError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar estado", "", err)
}
