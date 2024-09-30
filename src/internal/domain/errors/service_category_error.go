package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateCategoryError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar categoria", description, nil)
}
func UnknownFindCategoryError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar categoria", description, nil)
}
func UnknownDeleteCategoryError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar categoria", description, nil)
}
func UnknownUpdateCategoryError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar categoria", description, nil)
}
func ValidateCreateCategoryError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar categoria", "", err)
}