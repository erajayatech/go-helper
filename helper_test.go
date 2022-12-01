package helper

import (
	"reflect"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	type args struct {
		payload             interface{}
		payloadMessageError map[string]string
	}

	type request struct {
		Email string `json:"email" example:"andrietrilaksono@gmail.com" validate:"required" msg:"error_invalid_email"`
		Name  string `json:"name" example:"andrie" validate:"required" msg:"error_invalid_name"`
	}

	var requestErrorMessage = map[string]string{
		"error_invalid_email": "email is required",
		"error_invalid_name":  "name is required",
	}

	tests := []struct {
		name           string
		args           args
		wantErrMessage map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "first test validate struct",
			args: args{
				payload:             request{},
				payloadMessageError: requestErrorMessage,
			},
			wantErrMessage: map[string]string{
				"email": "email is required",
				"name":  "name is required",
			},
		},
		{
			name: "second test validate struct",
			args: args{
				payload: request{
					Name: "Andrie",
				},
				payloadMessageError: requestErrorMessage,
			},
			wantErrMessage: map[string]string{
				"email": "email is required",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrMessage := ValidateStruct(tt.args.payload, tt.args.payloadMessageError); !reflect.DeepEqual(gotErrMessage, tt.wantErrMessage) {
				t.Errorf("ValidateStruct() = %v, want %v", gotErrMessage, tt.wantErrMessage)
			}
		})
	}
}

func TestValidateStructWithError(t *testing.T) {
	type args struct {
		payload             interface{}
		payloadMessageError map[string]string
	}

	type request struct {
		Email string `json:"email" example:"andrietrilaksono@gmail.com" validate:"required" msg:"error_invalid_email"`
		Name  string `json:"name" example:"andrie" validate:"required" msg:"error_invalid_name"`
	}

	var requestErrorMessage = map[string]string{
		"error_invalid_email": "email is required",
		"error_invalid_name":  "name is required",
	}

	tests := []struct {
		name           string
		args           args
		wantErrMessage map[string]string
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name: "first test validate struct with error",
			args: args{
				payload:             request{},
				payloadMessageError: requestErrorMessage,
			},
			wantErrMessage: map[string]string{
				"email": "email is required",
				"name":  "name is required",
			},
			wantErr: true,
		},
		{
			name: "second test validate struct with error",
			args: args{
				payload: request{
					Name: "Andrie",
				},
				payloadMessageError: requestErrorMessage,
			},
			wantErrMessage: map[string]string{
				"email": "email is required",
			},
			wantErr: true,
		},
		{
			name: "third test validate struct with error",
			args: args{
				payload: request{
					Name:  "Andrie",
					Email: "andrietrilaksono@gmail.com",
				},
				payloadMessageError: requestErrorMessage,
			},
			wantErrMessage: map[string]string(nil),
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErrMessage, err := ValidateStructWithError(tt.args.payload, tt.args.payloadMessageError)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateStructWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotErrMessage, tt.wantErrMessage) {
				t.Errorf("ValidateStructWithError() = %v, want %v", gotErrMessage, tt.wantErrMessage)
			}
		})
	}
}

func TestFormatRupiah(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "first test rupiah",
			args: args{
				amount: 3000,
			},
			want: "Rp 3.000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatRupiah(tt.args.amount); got != tt.want {
				t.Errorf("FormatRupiah() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatGender(t *testing.T) {
	type args struct {
		gender int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test male",
			args: args{
				gender: 1,
			},
			want: "M",
		},
		{
			name: "test female",
			args: args{
				gender: 2,
			},
			want: "F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatGender(tt.args.gender); got != tt.want {
				t.Errorf("FormatGender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "first test must get env",
			args: args{
				key: "MODE",
			},
			want: "local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustGetEnv(tt.args.key); got != tt.want {
				t.Errorf("MustGetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInfoText(t *testing.T) {
	type args struct {
		actionName  string
		orderNumber string
		status      string
		updatedBy   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "first test format info text",
			args: args{
				actionName:  "update data",
				orderNumber: "11000034",
				status:      "processing",
				updatedBy:   "system",
			},
			want: "update data #11000034 processing - updated_by: system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatInfoText(tt.args.actionName, tt.args.orderNumber, tt.args.status, tt.args.updatedBy); got != tt.want {
				t.Errorf("FormatInfoText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpectedInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first test expected int",
			args: args{
				v: 6.5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpectedInt(tt.args.v); got != tt.want {
				t.Errorf("ExpectedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpectedInt64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "first test expected int64",
			args: args{
				v: 67,
			},
			want: 67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpectedInt64(tt.args.v); got != tt.want {
				t.Errorf("ExpectedInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpectedString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "first test expected string",
			args: args{
				v: 64,
			},
			want: "64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpectedString(tt.args.v); got != tt.want {
				t.Errorf("ExpectedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "first test float to string",
			args: args{
				f: 0.06,
			},
			want: "0.060000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(tt.args.f); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDateFormat(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "first test validate date format",
			args: args{
				p: "01-09-2010",
			},
			wantResult: "01-09-2010",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ValidateDateFormat(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDateFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ValidateDateFormat() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
