package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateUserError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar dados do usuario", description, nil)
}
func UnknownFindUserError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar dados do usuario", description, nil)
}
func NotFoundServiceProviderError() error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar dados do usuario",  "Provedor não encontrado.", nil)
}
func NotFoundUserError() error {
	return error_common.NewCustomError(error_common.ERR_USER_NOTFOUND,
		"Falha ao buscar dados do usuario",  "Usuario não encontrado.", nil)
}
func UnknownDeleteUserError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar dados do usuario", description, nil)
}
func UnknownAddFcmTokenUserError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao adicionar token fcm ao usuario", description, nil)
}
func UnknownUpdateUserError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar dados do usuario", description, nil)
}
func ValidateCreateUserError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar dados do usuario", "", err)
}
func InvalidCredentialError() error {
	return error_common.NewCustomError(error_common.ERR_INVALID_CREDENCIAS,
		"Falha ao buscar token de acesso", "Email ou senha estão incorretos.", nil)
}