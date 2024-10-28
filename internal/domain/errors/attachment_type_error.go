package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateAttachmentTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar tipo de anexo", description, nil)
}
func UnknownFindAttachmentTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar tipo de anexo", description, nil)
}
func UnknownDeleteAttachmentTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar tipo de anexo", description, nil)
}
func UnknownUpdateAttachmentTypeError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar tipo de anexo", description, nil)
}
func ValidateCreateAttachmentTypeError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar tipo de anexo", "", err)
}