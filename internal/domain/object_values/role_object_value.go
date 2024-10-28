package object_values

import "github.com/google/uuid"

const (
	ROLE_CLIENT_NAME           = "Cliente"
	ROLE_ADMIN_NAME            = "admin"
	ROLE_SERVICE_PROVIDER_NAME = "Prestador de serviço"
)

var (
	ROLE_CLIENT_ID          , _ = uuid.Parse("65c95ce2-1f2c-4325-a9a7-71af6d127cd9")
	ROLE_ADMIN_ID           , _ = uuid.Parse("ba162896-788c-486d-9d8f-e6f4ac334275")
	ROLE_SERVICE_PROVIDER_ID, _ = uuid.Parse("939e8771-5f21-47e5-85bc-bec7f9ae8848")
)
