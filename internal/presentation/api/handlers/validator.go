package handlers

import (
	"devflow/internal/presentation/api/responses"
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func buildValidationCauses(err error) []responses.Cause {
	out := make([]responses.Cause, 0)
	if err == nil {
		return out
	}
	var verrs validator.ValidationErrors
	if errors.As(err, &verrs) {
		for _, fe := range verrs {
			out = append(out, responses.Cause{
				Field:   fe.Field(),
				Message: fe.Error(),
			})
		}
		return out
	}
	out = append(out, responses.Cause{Message: err.Error()})
	return out
}
