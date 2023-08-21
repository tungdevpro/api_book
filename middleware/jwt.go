package middleware

import (
	"api_book/config"
	"api_book/helpers"
	"api_book/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenToken(user models.User) (string, error) {
	secretKey := config.Get("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("")
	}

	if len(user.ID) == 0 {
		return "", errors.New("len(id) == 0")
	}

	claims := &models.JwtCustomClaims{
		UserId: "1331231",
		Role:   user.Role,
		Phone:  user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3600).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(secretKey))
	helpers.Fatal(err)

	return result, nil
}

func DecodeToken() {}
