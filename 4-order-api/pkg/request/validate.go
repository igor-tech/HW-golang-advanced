package request

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func IsValid[T any](data T) error {
	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
