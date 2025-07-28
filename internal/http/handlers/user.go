package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rahulvarma07/note_backend/internal/config"
	"github.com/rahulvarma07/note_backend/internal/http/mail"
	"github.com/rahulvarma07/note_backend/internal/http/models"
	"github.com/rahulvarma07/note_backend/internal/http/utils"
)

func SendVerificationMail(successMail *config.Mail) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var userModel models.UserSignUp
		err := json.NewDecoder(r.Body).Decode(&userModel)

		// checking errors like End of file
		if errors.Is(err, io.EOF) {
			utils.SetResponse(w, http.StatusBadRequest, utils.GeneralErrors(err))
			return
		}
		// general error check
		if err != nil {
			utils.SetResponse(w, http.StatusBadRequest, utils.GeneralErrors(err))
			return
		}

		// checking for validation errors
		err = validator.New().Struct(userModel)
		if err != nil {
			valiadtionError := err.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(valiadtionError))
			return
		}

		// now if there is no error 
		// generate a mail to the user

		go mail.SendMail(successMail, &userModel)
		
		utils.SetResponse(w, http.StatusCreated, map[string]string{"message" : "verify email and login"})
		
	}
}

