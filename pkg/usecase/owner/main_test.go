// +build dev test

package owner

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// using a test table
var TestItems = []struct {
	comment       string // a comment used to identify test in output
	session       string // in our mock we use this value to affect the return values
	expectedError error
	expectedOwner string
}{
	{
		comment:       "send valid session",
		session:       "1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn",
		expectedError: nil,
		expectedOwner: "valid@session.owner",
	},
	{
		comment:       "trigger datastore error",
		session:       "111111lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn",
		expectedError: errors.New("datastore error"),
		expectedOwner: "",
	},
}

func TestRequestLink(t *testing.T) {

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	for _, item := range TestItems {

		owner, err := s.SessionOwner(item.session)
		assert.Equal(t, item.expectedOwner, owner, item.comment)
		assert.Equal(t, item.expectedError, err, item.comment)
	}
}
