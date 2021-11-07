// +build dev test

package auth

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/thisdougb/magiclink/config"
	"testing"
)

// using a test table
var TestItems = []struct {
	comment       string // a comment used to identify test in output
	magiclinkid   string // in our mock we use this value to affect the return values
	expectedError error
}{
	{
		comment:       "valid magiclinkid",
		magiclinkid:   "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		expectedError: nil,
	},
	{
		comment:       "too short magiclinkid",
		magiclinkid:   "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbp",
		expectedError: errors.New("invalid magic link id"),
	},
	{
		comment:       "too long magiclinkid",
		magiclinkid:   "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpbp",
		expectedError: errors.New("invalid magic link id"),
	},
	{
		comment:       "invalid char in magiclinkid",
		magiclinkid:   "M;m7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		expectedError: errors.New("invalid magic link id"),
	},
	{
		comment:       "magic link datastore error",
		magiclinkid:   "sVm4ECyEaec1HYBI9yP8nqLPMP1f8PXSar2O1ZN5HzyNn1WCr5Zx7JuInMUB8o8t",
		expectedError: errors.New("datastore error"),
	},
	{
		comment:       "magic link not found",
		magiclinkid:   "MBm7ECyEaec1HYBI9yP8nqLPMP1f8PXSar2O1ZN5HzyNn1WCr5Zx7JuInMUB8o8t",
		expectedError: errors.New("magic link not found"),
	},
}

func TestMagicLink(t *testing.T) {

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	for _, item := range TestItems {

		_, err := s.Auth(item.magiclinkid)
		assert.Equal(t, item.expectedError, err, item.comment)
	}
}

// using a test table
var TestSessionItems = []struct {
	comment       string // a comment used to identify test in output
	magiclinkid   string // in our mock we use this value to affect the return values
	expectedError error
}{
	{
		comment:       "valid magiclinkid",
		magiclinkid:   "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		expectedError: nil,
	},
	{
		comment:       "session id datastore error",
		magiclinkid:   "sessioniddatastoreerrordyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		expectedError: errors.New("datastore error"),
	},
}

func TestSessionID(t *testing.T) {

	var cfg *config.Config // dynamic config settings

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	for _, item := range TestSessionItems {

		sessionID, err := s.Auth(item.magiclinkid)
		if err != nil {
			assert.Equal(t, item.expectedError, err, item.comment)
			assert.Equal(t, 0, len(sessionID), item.comment)
		} else {
			assert.Equal(t, cfg.SESSION_ID_LENGTH(), len(sessionID), item.comment)
		}

	}
}
