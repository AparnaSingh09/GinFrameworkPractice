package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: getSecret(),
		issuer:    "thoughtworks.com",
	}
}

func getSecret() string {
	return "secret"
}

func (jwtService *JWTService) GenerateToken(userName string) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userName"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["Issuer"] = jwtService.issuer

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	tokenStr, _ := token.SignedString([]byte(jwtService.secretKey))
	return tokenStr
}

func (jwtService *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtService.secretKey), nil
	})
}
