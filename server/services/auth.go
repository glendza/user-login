package services

import (
	"glendza/login-code-challenge/config"
	"glendza/login-code-challenge/models"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	secretKey := []byte(config.GetConfig().SecretKey)
	claims := JwtClaims{
		UserID: user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseJWT(tokenString string) (int, error) {
	secretKey := []byte(config.GetConfig().SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims.UserID, err
	}

	return 0, err
}
