package models_responses

type TokenResponse struct {
	 AccessToken    string `json:"accessToken"`
	 ExpirationTime int64  `json:"expirationTime"`
}
