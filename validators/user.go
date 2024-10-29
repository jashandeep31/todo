package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type UserInput struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ValidateInput[T any](input T) (map[string]string, error) {
	if err := validate.Struct(input); err != nil {
		errs := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errs[err.Field()] = fmt.Sprintf("failed validation: %s", err.Tag())
		}
		return errs, fmt.Errorf("validation failed")
	}
	return nil, nil
}
