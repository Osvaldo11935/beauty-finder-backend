package object_values

import "github.com/google/uuid"

const (
	STATUS_CANCELLED_NAME = "Cancelado"
	STATUS_PENDING_NAME   = "Pendente"
	STATUS_COMPLETED_NAME = "Concluido"
)

var (
	STATUS_CANCELLED_ID = uuid.MustParse("03155a99-1ae7-4fa7-b55d-1a00e8f16a76")
	STATUS_PENDING_ID   = uuid.MustParse("39ab06e0-1530-4c51-ad84-4843d9978425")
	STATUS_COMPLETED_ID = uuid.MustParse("14f81cb0-9e2e-4556-a5bb-defc0c5efb4d")
)
