package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type CustomValidatorService struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidatorService {
	return &CustomValidatorService{}
}

func (cv *CustomValidatorService) Validate(i interface{}) error {
  log.Println("could not validate data")
	return cv.Validator.Struct(i)
}
