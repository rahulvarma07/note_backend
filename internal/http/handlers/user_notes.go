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

func AddNotes(notesCollection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer r.Body.Close()

		// decode the resposne
		var userNoteRequest models.UserNotesRequest
		if err := json.NewDecoder(r.Body).Decode(&userNoteRequest); err != nil {
			// set the response
			log.Println("unable to decode the request")
		}

		// valiade
		valErr := validator.New().Struct(userNoteRequest)
		if valErr != nil {
			err := valErr.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(err))
			return
		}

		// add the notes
		var noteDBModel models.UserNotesDataBase
		noteDBModel.UserId = userNoteRequest.UserId
		noteDBModel.NoteId = uuid.New().String()
		noteDBModel.Title = userNoteRequest.Title
		noteDBModel.Tag = userNoteRequest.Tag
		noteDBModel.Notes = userNoteRequest.Notes
		noteDBModel.CreatedTime = time.Now()
		noteDBModel.UpdatedTime = time.Now()

		_, err := notesCollection.InsertOne(ctx, noteDBModel)
		if err != nil {
			log.Println("Unable to add the note to the database", err)
		}

		utils.SetResponse(w, http.StatusCreated, "Successfully created note")
	}
}

func DeletUserNote(userNotesCollection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer r.Body.Close()

		var responseID models.UsertNotesRequestById
		if err := json.NewDecoder(r.Body).Decode(&responseID); err != nil {
			log.Println("unable to decode the response")
		}

		valErr := validator.New().Struct(responseID)
		if valErr != nil {
			err := valErr.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(err))
			return
		}

		// find the id and delete it
		filter := bson.M{"note_id": responseID.Id}
		_, err := userNotesCollection.DeleteOne(ctx, filter)
		if err != nil {
			utils.SetResponse(w, http.StatusBadRequest, utils.GeneralErrors(err))
			return
		}

		utils.SetResponse(w, http.StatusAccepted, map[string]string{"message": "successfully deleted note"})
	}
}


func GetAllNotes(usersNotesCollection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer  cancel()
		defer r.Body.Close()

		var requestId models.UsertNotesRequestById
		if err := json.NewDecoder(r.Body).Decode(&requestId); err != nil{
			log.Println("unable to decode")
		}

		valErr := validator.New().Struct(requestId)
		if valErr != nil{
			err := valErr.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(err))
			return
		}

		// user id with us
		// aggregation

		pipeline := mongo.Pipeline{
			{
				{Key: "$match", Value: bson.D{
					{Key : "user_id",Value:  requestId.Id},
				}},
			},
		}

		cursor, err := usersNotesCollection.Aggregate(ctx, pipeline)
		if err != nil{
			log.Println("aggregation error", err)
			return
		}

		var userNotes []bson.M
		if err := cursor.All(ctx, &userNotes); err != nil{
			log.Println("failed to read result")
			return
		}

		utils.SetResponse(w, http.StatusOK, userNotes)
	}
}

func GetNotesByTags(userNotesCollection *mongo.Collection) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer r.Body.Close()

		var userTagRequest models.UserNotesRequestByTag
		if err := json.NewDecoder(r.Body).Decode(&userTagRequest); err != nil{
			log.Println("unable to decode the body")
		}

		valErr := validator.New().Struct(userTagRequest)
		if valErr != nil{
			err := valErr.(validator.ValidationErrors)
			utils.SetResponse(w, http.StatusBadRequest, utils.CheckValidations(err))
			return
		}

		pipeline := mongo.Pipeline{
			{
				{Key: "$match", Value: bson.D{
					{Key: "user_id", Value: userTagRequest.Id},
				}},
			},
			{
				{Key: "$match", Value: bson.D{
					{Key: "tag", Value: userTagRequest.Tag},
				}},
			},
		}

		cursor, err := userNotesCollection.Aggregate(ctx, pipeline)
		if err != nil{
			log.Println("There is an error in aggregation")
			return 
		}

		var filterResult []bson.M
		if err := cursor.All(ctx, &filterResult); err != nil{
			log.Println("unable to read data after aggregation")
			return
		}

		utils.SetResponse(w, http.StatusOK, filterResult)
	}
}