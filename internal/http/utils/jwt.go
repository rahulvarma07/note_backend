package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rahulvarma07/note_backend/internal/http/models"
)

var secretKey  = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwtToken(user *models.UserSignUp)( string, error) {
	claims := jwt.MapClaims{
		"name" : user.Name,
		"email" : user.Email,
		"password" : user.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // using alfo HS256
	tokenString, err := token.SignedString(secretKey) 
	if err != nil{
		return "", nil
	}

	return tokenString, nil
}