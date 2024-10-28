package object_values

import "github.com/google/uuid"

const (
	RATING_TYPE_BAD_NAME = "Ruim"
	RATING_TYPE_NORMAL_NAME = "Normal"
	RATING_TYPE_EXCELLENT_NAME   = "Excelente"

)

var (
	RATING_TYPE_BAD_ID = uuid.MustParse("03155a99-1ae7-4fa7-b55d-1a00e8f16a76")
	RATING_TYPE_NORMAL_ID   = uuid.MustParse("39ab06e0-1530-4c51-ad84-4843d9978425")
	RATING_TYPE_EXCELLENT_ID = uuid.MustParse("14f81cb0-9e2e-4556-a5bb-defc0c5efb4d")
)
