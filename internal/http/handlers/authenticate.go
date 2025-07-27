package handlers

import "net/http"

func AuthenticateUser() http.HandlerFunc{

	// in this handler function
	// we will generate a JWT token for the user and store it in the database

	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}