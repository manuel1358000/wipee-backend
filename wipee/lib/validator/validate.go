package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct recibe cualquier estructura y valida sus campos basados en las etiquetas `validate`
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// ParseValidationErrors convierte los errores de validación en un mensaje de error más legible
func ParseValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var sb strings.Builder
		for _, fieldError := range validationErrors {
			sb.WriteString(fieldError.Field() + ": " + fieldError.Tag() + "; ")
		}
		return sb.String()
	}
	return "Validation failed"
}
