package jwt_handler

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(userAccountID string, secrectKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userAccountID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	secretKeyByte := []byte(secrectKey)
	tokenString, err := token.SignedString(secretKeyByte)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string, secrectKey string) (string, error) {
	secretKeyByte := []byte(secrectKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKeyByte, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return "", fmt.Errorf("token is expired")
	}

	return claims["sub"].(string), nil
}
