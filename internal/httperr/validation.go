package httperr

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func FieldErrorsFromBinding(err error) []FieldError {
	var ves validator.ValidationErrors
	if errors.As(err, &ves) {
		out := make([]FieldError, 0, len(ves))
		for _, fe := range ves {
			out = append(out, FieldError{
				Field: fe.Field(),
				Reason: fe.ActualTag(), // 例: "required", "email" など
			})
		}
		return out
	}
	return []FieldError{
		{
			Field:  "unknown",
			Reason: err.Error(),
		},
	}
}