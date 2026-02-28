package response

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error, req any) any {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["error"] = err.Error()
		return errors
	}

	t := reflect.TypeOf(req)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, e := range validationErrors {
		field, _ := t.FieldByName(e.StructField())

		jsonTag := field.Tag.Get("json")

		name := strings.Split(jsonTag, ",")[0]

		if name == "" || name == "-" {
			name = strings.ToLower(e.Field())
		}

		errors[name] = e.Tag()
	}

	return errors
}
