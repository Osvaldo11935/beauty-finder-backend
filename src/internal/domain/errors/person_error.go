package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreatePersonError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar dados pessoas", description, nil)
}
func UnknownFindPersonError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar dados pessoas", description, nil)
}
func UnknownDeletePersonError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar dados pessoas", description, nil)
}
func UnknownUpdatePersonError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar dados pessoas", description, nil)
}
func ValidateCreatePersonError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar dados pessoas", "", err)
}