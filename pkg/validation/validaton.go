package validation

import (
	"github.com/go-playground/validator/v10"
)

type RequestError struct {
	Field   string
	Message string
}

func MessageTag(tag string) string {
	switch tag {
	case "required":
		return "field is required"
	case "email":
		return "field must be a valid email"
	case "min":
		return "field must be at least 6 characters"
	case "number":
		return "field must be a number"
	case "gt":
		return "field must be greater than 8"
	case "url":
		return "field must be a valid URL"
	}
	return ""
}

func NewErrResponse(ve validator.ValidationErrors) []RequestError {
	out := make([]RequestError, len(ve))
	for i, fe := range ve {
		out[i] = RequestError{
			Field:   fe.Field(),
			Message: MessageTag(fe.Tag()),
		}
	}
	return out

}
