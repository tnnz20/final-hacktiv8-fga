package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type RequestError struct {
	Field   string
	Message string
}

func MessageTag(tag string) string {
	switch tag {
	case "required":
		return "cannot be empty"
	case "email":
		return "must be a valid email address"
	case "min":
		return "must be at least 6 characters"
	case "number":
		return "must be a number"
	case "gt":
		return "must be greater than 8"
	case "url":
		return "must be a valid URL"
	}
	return ""
}

func NewErrResponse(ve validator.ValidationErrors) string {
	out := make([]RequestError, len(ve))
	for i, fe := range ve {
		out[i] = RequestError{
			Field:   fe.Field(),
			Message: MessageTag(fe.Tag()),
		}
	}
	return fmt.Sprintf("error: field %s %s", out[0].Field, out[0].Message)

}
