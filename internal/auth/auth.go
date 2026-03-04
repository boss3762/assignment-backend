package auth

import (
	"github.com/golang-jwt/jwt"
	"time"
	"fmt"
)

type AuthClaims struct{
	Username string
	jwt.StandardClaims
}

type JWTService interface {
	GenerateToken(username string) string
	ValidateToken(token string) (string, error)
}

func NewAuthService() JWTService {
	return &AuthClaims{}
}

func (a *AuthClaims) GenerateToken(username string) string {
	claims := AuthClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}
	return tokenString
}

func (a *AuthClaims) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	if claims, ok := token.Claims.(*AuthClaims); ok {
		return claims.Username, nil
	}

	return "", fmt.Errorf("invalid token")
}