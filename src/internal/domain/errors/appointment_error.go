package errors

import error_common "src/internal/domain/errors/common"

func UnknownCreateAppointmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao cadastrar agendamento", description, nil)
}
func UnknownFindAppointmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao buscar agendamento", description, nil)
}
func UnknownDeleteAppointmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao deletar agendamento", description, nil)
}
func UnknownUpdateAppointmentError(description string) error {
	return error_common.NewCustomError(error_common.ERR_UNKNOWN,
		"Falha ao atualizar agendamento", description, nil)
}
func ValidateCreateAppointmentError(err []string) error {
	return error_common.NewCustomError(error_common.ERR_VALIDATE,
		"Ocorreu um erro ao registrar agendamento", "", err)
}