package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rahulvarma07/note_backend/internal/http/models"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// this function is generate token for mail
func MailToken(user *models.UserSignUp) (string, error) {
	claims := jwt.MapClaims{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // using alfo HS256
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

// this function is to generate the login token
func UserAuthToken(user *models.UserDataBaseModel) (string, error) {
	claims := jwt.MapClaims{
		"user_id":       user.UserId,
		"user_name":     user.UserName,
		"user_email":    user.UserEmail,
		"user_password": user.UserPassword,
	}
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return token, err
}

// this function is to parse the JWT token..
func GetTokenInfo(token string) (*models.UserSignUp, error) {

	claims := jwt.MapClaims{}

	// parsing with secretkey
	tokenString, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !tokenString.Valid {
		return nil, err
	}

	return &models.UserSignUp{
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		Password: claims["password"].(string),
	}, nil
}
