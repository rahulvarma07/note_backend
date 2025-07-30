package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rahulvarma07/note_backend/internal/config"
	"github.com/rahulvarma07/note_backend/internal/http/mail"
	"github.com/rahulvarma07/note_backend/internal/http/models"
	"github.com/rahulvarma07/note_backend/internal/http/utils"
	"github.com/rahulvarma07/note_backend/internal/messages"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UserVerification(successMail *config.Mail, userAuthCollections *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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

		filter := bson.M{"user_email": userModel.Email}
		var isUserFound bson.M

		checkErr := userAuthCollections.FindOne(ctx, filter).Decode(&isUserFound)

		if checkErr != nil {
			if checkErr == mongo.ErrNoDocuments {
				// user not found..
				go mail.SendMail(successMail, &userModel)
				utils.SetResponse(w, http.StatusCreated, map[string]string{"message": messages.SuccessMail})
			} else {
				log.Fatal("There is an error connecting to database", err)
			}
			return
		}

		utils.SetResponse(w, http.StatusBadRequest, utils.CustomError(messages.MailExists))
	}
}
