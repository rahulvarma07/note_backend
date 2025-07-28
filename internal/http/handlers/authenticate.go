package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/rahulvarma07/note_backend/internal/http/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// this function is supposed to add the user to the database
func SignUpUser(mongoCollection *mongo.Collection) http.HandlerFunc {

	// in this handler function
	// we will generate a JWT token for the user and store it in the database

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		query := r.URL.Query().Get("token") // getting the token from backend
		userSignUpRequestDetails, err := utils.GetTokenInfo(query)

		if err != nil {
			// make resposnse writers
		}

		// check whether the user already exsits in the database..
		checkExisistingUser := bson.M{"mail": userSignUpRequestDetails.Email}
		isUserExsist := mongoCollection.FindOne(ctx, checkExisistingUser)

		if isUserExsist == nil {
			// make a response writer that already a user exsists
		}

		// if not
		// make the user stored inside the DB..
		
	}
}

// this function is supposed to login the user checking user credentials
func LoginTheUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
