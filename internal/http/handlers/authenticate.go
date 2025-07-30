package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rahulvarma07/note_backend/internal/http/models"
	"github.com/rahulvarma07/note_backend/internal/http/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// this function is supposed to add the user to the database
func SignUpUser(usersAuthCollection *mongo.Collection) http.HandlerFunc {

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

		_, dbErr := usersAuthCollection.InsertOne(ctx, storeUserModel)
		if dbErr != nil {
			log.Println("Error logging in a user", err)
		}
		utils.SetResponse(w, http.StatusCreated, map[string]string{"message": "Verification successfull, login with the credentials"})
	}
}

// this function is supposed to login the user checking user credentials
func LoginTheUser(usersAuthCollection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer r.Body.Close()

		// decode the request body
		var userDetails models.UserLogin
		if err := json.NewDecoder(r.Body).Decode(&userDetails); err != nil {
			log.Println("Unable to decode the request")
		}
		// validating the decoded data
		valErr := validator.New().Struct(userDetails)
		if valErr != nil {
			err := valErr.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(err))
			return 
		}
		// check whether the email is present
		filter := bson.M{"user_email": userDetails.Email}
		var userDBModel models.UserDataBaseModel
		findTheUser := usersAuthCollection.FindOne(ctx, filter).Decode(&userDBModel)
		if findTheUser != nil{
			if findTheUser == mongo.ErrNoDocuments{
				utils.SetResponse(w, http.StatusUnauthorized, utils.CustomError("Invalid authentication credentials"))
			}else{
				log.Fatal("Database error ", findTheUser)
			}
			return
		}
		
		// check for passwords matching
		userEnteredPassword := userDetails.Password
		userDBHashedPassword := userDBModel.UserPassword
		var passwordsMatch bool = utils.CheckPasswords(userEnteredPassword, userDBHashedPassword)
		if !passwordsMatch{
			utils.SetResponse(w, http.StatusUnauthorized, utils.CustomError("Invalid authentication credentials"))
			return 
		}
		
		// finally generate the token for user as response
		userToken, err := utils.UserAuthToken(&userDBModel)
		if err != nil{
			log.Fatal("Unable to generate the token", err)
		}

		utils.SetResponse(w, http.StatusAccepted, map[string]string{"message" : userToken})
	}
}
