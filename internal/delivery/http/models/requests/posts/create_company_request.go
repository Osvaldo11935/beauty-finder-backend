package models_requests_posts

type CreateCompanyRequest struct {
	Name             string `json:"name"`
	PhoneNumber      string `json:"phoneNumber"`
	NationalRegistry string `json:"nationalRegistry"`
	EmployeeNumber   int    `json:"employeeNumber"`
}
