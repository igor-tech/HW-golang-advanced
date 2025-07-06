package request

import "github.com/go-playground/validator/v10"

func IsValid[T any](data T) error {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
