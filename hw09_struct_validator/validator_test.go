package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "sample string",
			expectedErr: ErrUnsupportedInterfaceValue,
		},
		{
			in: User{
				ID:     "string must be 36 symbols of     len",
				Name:   "Alexey",
				Age:    36,
				Email:  "alx@mail.ru",
				Role:   "admin",
				Phones: []string{"89991112244", "87776665533"},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "failed string",
				Name:   "Alexey",
				Age:    12,
				Email:  "failed.mail.ru",
				Role:   "notadmin",
				Phones: []string{"89991112244", "87776665533"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   ErrStringLen,
				},
				ValidationError{
					Field: "Age",
					Err:   ErrNumberMin,
				},
				ValidationError{
					Field: "Email",
					Err:   ErrStringRegex,
				},
				ValidationError{
					Field: "Role",
					Err:   ErrStringSet,
				},
			},
		},
		{
			in: App{
				Version: "1.0.1.2",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   ErrStringLen,
				},
			},
		},
		{
			in: Response{
				Code: 502,
				Body: "",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   ErrNumberSet,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
