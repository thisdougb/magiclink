// +build dev test

package auth

import (
	"errors"
)

type MockWriter struct{}

//
// Mock methods here, with conditionals enabling testing of return values
//
func (m *MockReader) GetExpireAccountFromID(magiclinkid string) (string, error) {

	if magiclinkid == "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb" {
		return "someuser@domain.com", nil
	}
	if magiclinkid == "sVm4ECyEaec1HYBI9yP8nqLPMP1f8PXSar2O1ZN5HzyNn1WCr5Zx7JuInMUB8o8t" {
		return "", errors.New("datastore error")
	}

	if magiclinkid == "sessioniddatastoreerrordyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb" {
		return "sessionid@datastore.error", nil
	}

	return "", nil
}

func (m *MockReader) StoreSessionID(email string, sessionID string, ttlSeconds int) error {

	if email == "someuser@domain.com" {
		return nil
	}
	if email == "sessionid@datastore.error" {
		return errors.New("datastore error")
	}

	return nil
}
