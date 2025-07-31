package models

import "time"

type UserNotesDataBase struct {
	UserId      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Tag         string    `json:"tag" bson:"tag"`
	Notes       string    `json:"note" bson:"note"`
	CreatedTime time.Time `json:"created_time" bson:"created_time"`
	UpdatedTime time.Time `json:"updated_time" bson:"updated_time"`
}

type UserNotesRequest struct {
	UserId string `validate:"required" json:"user_id" bson:"user_id"`
	Title  string `validate:"required" json:"title" bson:"title"`
	Tag    string `validate:"required" json:"tag" bson:"tag"`
	Notes  string `validate:"required" json:"note" bson:"note"`
}
