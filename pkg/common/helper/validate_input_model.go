package helper

import (
	"github.com/go-playground/validator/v10"
)

type SchemaFiledErr struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateInputModel[T any](schema T) (schemaErrs []SchemaFiledErr, err error) {
	e := Validator.Struct(schema)

	if e != nil {
		for _, err := range e.(validator.ValidationErrors) {
			var schemaFileErr SchemaFiledErr
			schemaFileErr.Field = err.Field()
			schemaFileErr.Tag = err.Tag()
			schemaFileErr.Value = err.Param()
			schemaErrs = append(schemaErrs, schemaFileErr)
		}

		return schemaErrs, e
	}

	return schemaErrs, nil
}
