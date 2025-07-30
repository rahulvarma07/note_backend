package utils

import "golang.org/x/crypto/bcrypt"

// function to hash the password
func HashPassword(userPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// function to check whether the hashed password and user entered password are same
func CheckPasswords(userEnteredPassword, DatabaseStoredPassword string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(DatabaseStoredPassword), []byte(userEnteredPassword))
	return err == nil
}