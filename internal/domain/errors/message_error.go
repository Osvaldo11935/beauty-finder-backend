package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateMessageError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar mensagem", description, nil)
}
func UnknownFindMessageError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar mensagem", description, nil)
}
func UnknownDeleteMessageError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar mensagem", description, nil)
}
func UnknownUpdateMessageError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar mensagem", description, nil)
}
func ValidateCreateMessageError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar mensagem", "", err)
}