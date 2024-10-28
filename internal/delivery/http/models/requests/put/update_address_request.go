package models_requests_puts


type UpdateAddressRequest struct  {
	 District     *string  `json:"district"`
	 Commune      *string  `json:"commune"`
	 Province     *string  `json:"province"`
	 Country      *string  `json:"country"`
	 City         *string  `json:"city"`
	 Street       *string  `json:"street"`
	 Neighborhood *string  `json:"neighborhood"`
	 Latitude     *float64 `json:"latitude"`
	 Longitude    *float64 `json:"longitude"`
}