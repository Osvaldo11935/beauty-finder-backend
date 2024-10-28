package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateRatingTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar tipo de avaliação", description, nil)
}
func UnknownFindRatingTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar tipo de avaliação", description, nil)
}
func UnknownDeleteRatingTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar tipo de avaliação", description, nil)
}
func UnknownUpdateRatingTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar tipo de avaliação", description, nil)
}
func ValidateCreateRatingTypeError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar tipo de avaliação", "", err)
}
