// +build dev test

package send

import (
	"errors"
	"strings"
)

type MockWriter struct{}

//
// Mock methods here, with conditionals enabling testing of return values
//
func (m *MockWriter) SubmitSendLinkRequest(data string) error {

	if strings.Contains(data, "fail@datastore.error") {
		return errors.New("datastore error")
	}

	return nil
}

func (m *MockWriter) StoreAuthID(email string, id string, ttlSeconds int) error {
	if email == "StoreAuthID@datastore.error" {
		return errors.New("datastore error")
	}
	return nil
}
