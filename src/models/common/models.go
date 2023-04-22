package common

import "github.com/go-playground/validator/v10"

type User struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username" validate:"required"`
	UserPassword string `json:"user_password" validate:"required"`
	UserEmail    string `json:"user_email" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
}

var validate *validator.Validate

func (u User) CheckRegisterRequest() bool {

	validate = validator.New()

	err := validate.Struct(u)

	if err != nil {
		return true
	} else {
		return false
	}
}
