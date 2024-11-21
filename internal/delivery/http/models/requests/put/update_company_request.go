package models_requests_puts

type UpdateCompanyRequest struct {
	Name             *string `json:"name"`
	PhoneNumber      *string `json:"phoneNumber"`
	NationalRegistry *string `json:"nationalRegistry"`
	EmployeeNumber   *int    `json:"employeeNumber"`
}
