package services

import (
	"errors"
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

func ParseJWT(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(config.GetConfig().SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		err := errors.New("invalid token")
		return nil, err
	}

	return token, nil
}

func ExtractUserId(token *jwt.Token) int {
	claims, _ := token.Claims.(*JwtClaims)
	return claims.UserID
}
