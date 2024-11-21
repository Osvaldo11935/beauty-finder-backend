package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateCompanyError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar empresa", description, nil)
}
func UnknownFindCompanyError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar empresa", description, nil)
}
func NotFoundFindCompanyError() error {
	return error_common.NewCustomError(error_common.ERR_NOTFOUND,
		"Falha ao buscar empresa", "Empresa não encontrada.", nil)
}
func UnknownDeleteCompanyError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar empresa", description, nil)
}
func UnknownUpdateCompanyError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar empresa", description, nil)
}
func ValidateCreateCompanyError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar empresa", "", err)
}
