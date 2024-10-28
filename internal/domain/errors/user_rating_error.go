package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateUserRatingError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar avaliação do usuario", description, nil)
}
func UnknownFindUserRatingError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar avaliação do usuario", description, nil)
}
func NotFoundFindUserRatingError() error {
	return error_common.NewCustomError(error_common.ERR_NOTFOUND,
		"Falha ao buscar avaliação do usuario", "Avaliação do usuario não encontrado", nil)
}
func UnknownDeleteUserRatingError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar avaliação do usuario", description, nil)
}
func UnknownUpdateUserRatingError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar avaliação do usuario", description, nil)
}
func ValidateCreateUserRatingError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar avaliação do usuario", "", err)
}