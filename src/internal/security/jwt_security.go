package security

import (
	"errors"
	"log"
	"src/internal/configs"
	models_responses "src/internal/delivery/http/models/responses"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ClaimJwt struct {
	Email  string   `json:"email"`
	UserID string   `json:"userId"`
	Claims []string `json:"claims"`
	jwt.StandardClaims
}

type JwtTokenService struct {
	JwtKey []byte
}

func NewJwtTokenService() *JwtTokenService {
	config, configErr := configs.LoadConfig()

	if configErr != nil {
		log.Print("Ocorreu um erro ao buscar configurações jwt")
		return nil
	}

	return &JwtTokenService{
		JwtKey: []byte(config.JwtSecretKey),
	}
}

func (s *JwtTokenService) GenerateToken(email *string, userId string, claim []string) (*models_responses.TokenResponse, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &ClaimJwt{
		UserID: userId,
		Claims: claim,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	if email != nil {
		claims.Email = *email
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.JwtKey)
	if err != nil {
		return nil, err
	}

	return &models_responses.TokenResponse{
		AccessToken:    tokenString,
		ExpirationTime: expirationTime.Unix(),
	}, nil
}

func (s *JwtTokenService) ValidateToken(tokenString string) (*ClaimJwt, error) {
	claims := ClaimJwt{}
	tokenString = strings.TrimSpace(tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return s.JwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return &claims, nil
}
