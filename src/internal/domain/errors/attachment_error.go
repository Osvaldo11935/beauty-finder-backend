package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateAttachmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar anexo", description, nil)
}
func UnknownFindAttachmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar anexo", description, nil)
}
func UnknownDeleteAttachmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar anexo", description, nil)
}
func UnknownUpdateAttachmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar anexo", description, nil)
}
func ValidateCreateAttachmentError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar anexo", "", err)
}