package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ConvertInterface(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

func TestValidateStruct(t *testing.T) {
	type request struct {
		Email string `json:"email" example:"andrietrilaksono@gmail.com" validate:"required" msg:"error_invalid_email"`
		Name  string `json:"name" example:"andrie" validate:"required" msg:"error_invalid_name"`
	}

	var requestErrorMessage = map[string]string{
		"error_invalid_email": "email is required",
		"error_invalid_name":  "name is required",
	}

	type test struct {
		payload request
		want    map[string]string
	}

	tests := []test{
		{
			payload: request{
				Name: "andrie",
			},
			want: map[string]string{
				"email": "email is required",
			},
		},
		{
			payload: request{
				Email: "andrietrilaksono@gmail.com",
			},
			want: map[string]string{
				"name": "name is required",
			},
		},
		{
			payload: request{},
			want: map[string]string{
				"name":  "name is required",
				"email": "email is required",
			},
		},
	}

	for _, testcase := range tests {
		assert.Equal(t, testcase.want, ValidateStruct(testcase.payload, requestErrorMessage))
	}
}

func TestValidateStructWithError(t *testing.T) {
	type request struct {
		Email string `json:"email" example:"andrietrilaksono@gmail.com" validate:"required" msg:"error_invalid_email"`
		Name  string `json:"name" example:"andrie" validate:"required" msg:"error_invalid_name"`
	}

	var requestErrorMessage = map[string]string{
		"error_invalid_email": "email is required",
		"error_invalid_name":  "name is required",
	}

	type response struct {
		data map[string]string
		err  error
	}

	type test struct {
		payload request
		want    response
	}

	tests := []test{
		{
			payload: request{
				Name:  "andrie",
				Email: "andrietrilaksono@gmail.com",
			},
			want: response{
				data: map[string]string(nil),
				err:  nil,
			},
		},
	}

	for _, testcase := range tests {
		assert.Equal(t, ConvertInterface(testcase.want.data, testcase.want.err), ConvertInterface(ValidateStructWithError(testcase.payload, requestErrorMessage)))
	}
}
