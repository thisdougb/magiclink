// +build dev test

package owner

import (
	"errors"
)

type MockReader struct {
}

func (m *MockReader) GetSessionOwner(session string) (string, error) {

	if session == "1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn" {
		return "valid@session.owner", nil
	}
	if session == "111111lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn" {
		return "", errors.New("datastore error")
	}
	return "", nil
}
