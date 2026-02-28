package middleware

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("json"), ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &CustomValidator{
		validator: v,
	}
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}
