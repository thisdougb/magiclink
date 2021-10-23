// +build dev test

package send

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// using a test table
var TestItems = []struct {
	comment       string // a comment used to identify test in output
	email         string // in our mock we use this value to affect the return values
	expectedError error
}{
	{
		comment:       "send valid email",
		email:         "user@domain.com",
		expectedError: nil,
	},
	{
		comment:       "invalid email",
		email:         "@domain.com",
		expectedError: errors.New("email invalid"),
	},
	{
		comment:       "datastore error",
		email:         "fail@datastore.error",
		expectedError: errors.New("datastore error"),
	},
	{
		comment:       "StoreAuthID datastore error",
		email:         "StoreAuthID@datastore.error",
		expectedError: errors.New("datastore error"),
	},
}

func TestRequestLink(t *testing.T) {

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	for _, item := range TestItems {

		err := s.Send(item.email)
		assert.Equal(t, item.expectedError, err, item.comment)
	}
}
