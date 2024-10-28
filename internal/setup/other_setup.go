package setup

import (
	"src/internal/services"
	service_interface "src/internal/services/interface_services"
)

type OtherSetup struct {
	FileManager           service_interface.IFileManager
}

func NewOtherSetup() *OtherSetup {
	return &OtherSetup{
		FileManager: services.NewGoogleDriveService(),
	}
}
