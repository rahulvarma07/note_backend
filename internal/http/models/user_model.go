package models

type UserLogin struct{
	Email string `validate:"required email" json:"user_email"`
	Password string `valiadet:"required" json:"user_password"`
}

type UserSignUp struct{
	Name string `validate:"required" json:"user_name"`
	Email string `validate:"email required" json:"user_email"`
	Password string `validate:"required" json:"user_password"`
}