package errors

import error_common "src/internal/domain/errors/common"

func ConfigurationFileNotFoundError() error {

	return error_common.NewCustomError(error_common.ERR_CONFIGURATION_FILE_NOT_FOUND,
		"Falha ao efetuar o upload do documento.", "Ficheiro de configuração não encontrado", nil)
}

func NewGoogleStorageUnknownError(description string) error {
	return error_common.NewCustomError(error_common.ERR_GOOGLE_STORAGE_UNKNOWN,
		"Falha ao efetuar o upload do documento.", description, nil)
}
