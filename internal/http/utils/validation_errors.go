package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// checks for various validation errors...

func CheckValidations(err validator.ValidationErrors) SetResponseModel {
	var valiadtionError []string
	for _, err := range err {
		switch err.ActualTag(){
		case "required":
			valiadtionError = append(valiadtionError, fmt.Sprintf("the field %s is required", err.Field()))
		case "email":
			valiadtionError = append(valiadtionError, fmt.Sprintf("the field %s is required", err.Field()))
		default:
			valiadtionError = append(valiadtionError, fmt.Sprintf("the field %s is required", err.Field()))
		}
	}
	
	return SetResponseModel{
		Status: statusFailure,
		Message: strings.Join(valiadtionError, ","),
	}
}
