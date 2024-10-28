package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateRoleError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar perfil", description, nil)
}
func UnknownFindRoleError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar perfil", description, nil)
}
func UnknownDeleteRoleError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar perfil", description, nil)
}
func UnknownUpdateRoleError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar perfil", description, nil)
}
func ValidateCreateRoleError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar perfil", "", err)
}