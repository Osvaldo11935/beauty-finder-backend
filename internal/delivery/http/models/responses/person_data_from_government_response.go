package models_responses

import (
	extensios "src/internal/domain/extensions"
	"github.com/google/uuid"
)

type PersonByNationalRegistry struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    PersonDataFromGovernmentResponse `json:"data"`
}

type PersonDataFromGovernmentResponse struct {
	Id                    uuid.UUID            `json:"id"`
	FullName              string               `json:"nome"`
	BirthDate             *extensios.CustomDate `json:"data_nasc"`
	Gender                string               `json:"genero"`
	Naturalness           string               `json:"naturalidade"`
	MaritalStatus         string               `json:"estado_civil"`
	FatherName            string               `json:"pai_nome_completo"`
	MotherName            string               `json:"mae_nome_completo"`
	NationalRegistry      string               `json:"nif"`
	PlaceIssuanceDocument string               `json:"emissao_local"`
	DateIssueDocument     *extensios.CustomDate `json:"data_emissao"`
}

