package handlers

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
	err := validate.RegisterValidation("customState", validState)
	if err != nil {
		panic(err)
	}
}
