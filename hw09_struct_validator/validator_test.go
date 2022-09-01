//nolint
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
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
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

	Order struct {
		Price []int `validate:"min:10"`
	}

	Bucket struct {
		Any string `validate:"qwerty:qwerty"`
	}

	Product struct {
		Name string `validate:"len:qwerty"`
	}

	Banana struct {
		Size int `validate:"min:qwerty"`
	}

	Ananas struct {
		Price int `validate:"max:qwerty"`
	}

	Car struct {
		Color string `validate:"regexp:^\\_$d++w+@\\w+\\.\\w+$"`
	}

	Bus struct {
		Any int `validate:"in:q,w,e,r,t,y"`
	}
)

func TestValidateFailedStruct(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{[]int{0}, ErrIsNotStruct},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := Validate(tt.in)
			require.Error(t, err)

			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{in: User{"f6b6fca6-7e4b-4966-bb7f-e8b531cdc109", "test", 20, "test@mail.ru", "admin", []string{"89181234567"}, []byte{}}},
		{in: App{Version: "12345"}},
		{in: Token{Header: []byte("test"), Payload: []byte("test"), Signature: []byte("test")}},
		{in: Response{Code: 200, Body: "test"}},
		{in: Order{Price: []int{10, 20, 30}}},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := Validate(tt.in)
			require.Nil(t, err)
		})
	}
}
