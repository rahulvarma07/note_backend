package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/rahulvarma07/note_backend/internal/http/models"
	"github.com/rahulvarma07/note_backend/internal/http/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// this function is supposed to add the user to the database
func SignUpUser(userAuthCollection *mongo.Collection) http.HandlerFunc {

	// in this handler function
	// we will generate a JWT token for the user and store it in the database

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer r.Body.Close()

		tokenString := r.URL.Query().Get("token")
		tokenInfo, err := utils.GetTokenInfo(tokenString)
		if err != nil {
			log.Println("unable to get the token information", err)
		}

		if err != nil {
			log.Println("Unable to hash the user password")
		}
		uniqueID := uuid.New().String()

		// make a struct which has user uuid, name, email and hashedPassword
		var storeUserModel models.UserDataBaseModel
		storeUserModel.UserId = uniqueID
		storeUserModel.UserName = tokenInfo.Name
		storeUserModel.UserEmail = tokenInfo.Email
		storeUserModel.UserPassword = tokenInfo.Password

		_, dbErr := userAuthCollection.InsertOne(ctx, storeUserModel)
		if dbErr != nil {
			log.Println("Error logging in a user", err)
		}
		utils.SetResponse(w, http.StatusCreated, map[string]string{"message": "Verification successfull, login with the credentials"})
	}
}

// this function is supposed to login the user checking user credentials
func LoginTheUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
