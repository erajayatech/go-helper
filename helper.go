package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/shyandsy/ShyGinErrors"
)

func ValidateStruct(payload interface{}, payloadMessageError map[string]string) (errMessage map[string]string) {
	ge := ShyGinErrors.NewShyGinErrors(payloadMessageError)
	err := CheckStruct(payload)
	if err != nil {
		errMessage = ge.ListAllErrors(payload, err)

		return
	}

	return
}

func CheckStruct(payload interface{}) (err error) {
	v := validator.New()

	return v.Struct(payload)
}
