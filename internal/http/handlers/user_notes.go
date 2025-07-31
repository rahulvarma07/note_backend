package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rahulvarma07/note_backend/internal/http/models"
	"github.com/rahulvarma07/note_backend/internal/http/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddNotes(notesCollection *mongo.Collection) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer r.Body.Close()

		// decode the resposne
		var userNoteRequest models.UserNotesRequest
		if err := json.NewDecoder(r.Body).Decode(&userNoteRequest); err != nil{
			// set the response
			log.Println("unable to decode the request")
		}

		// valiade 
		valErr := validator.New().Struct(userNoteRequest)
		if valErr != nil{
			err := valErr.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(err))
			return
		}

		// add the notes
		var noteDBModel models.UserNotesDataBase
		noteDBModel.UserId = userNoteRequest.UserId
		noteDBModel.Title = userNoteRequest.Title
		noteDBModel.Tag = userNoteRequest.Tag
		noteDBModel.Notes = userNoteRequest.Notes
		noteDBModel.CreatedTime = time.Now()
		noteDBModel.UpdatedTime = time.Now()

		_, err := notesCollection.InsertOne(ctx, noteDBModel)
		if err != nil{
			log.Println("Unable to add the note to the database", err)
		}

		utils.SetResponse(w, http.StatusCreated, "Successfully created note")
	}
}