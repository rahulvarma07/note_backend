package models

type UserDataBaseModel struct {
	UserId       string `bson:"user_id" json:"user_id"`
	UserName     string `bson:"user_name" json:"user_name"`
	UserEmail    string `bson:"user_email" json:"user_email"`
	UserPassword string `bson:"user_password" json:"user_password"`
}
