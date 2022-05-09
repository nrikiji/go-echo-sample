package error

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationMessages map[string]map[string]string

type ValidationError struct {
	Messages []string `json:"messages"`
}

func (e ValidationError) Error() string {
	return strings.Join(e.Messages, ",")
}

func NewValidationError(err error, m map[string]map[string]string) error {
	ve := err.(validator.ValidationErrors)
	if errors.As(err, &ve) {
		messages := []string{}
		for _, fe := range ve {
			messages = append(messages, MsgForTag(fe, m))
		}
		return &ValidationError{Messages: messages}
	}
	return nil
}

func MsgForTag(fe validator.FieldError, m map[string]map[string]string) string {
	if val, ok := m[fe.Field()][fe.Tag()]; ok {
		return val
	} else {
		return fe.Error()
	}
}
