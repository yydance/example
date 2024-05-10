package gvalidator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	XValidate = validator.New()
)

type XValidator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

func New() XValidator {
	return XValidator{
		validator: XValidate,
	}
}

func (v XValidator) Validate(data any) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	errs := XValidate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}

func (v XValidator) ErrMsgs(data any) []string {
	errs := v.Validate(data)
	errMsgs := make([]string, 0)
	if len(errs) > 0 && errs[0].Error {
		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: %v Needs to implement %s",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
	}
	return errMsgs
}
