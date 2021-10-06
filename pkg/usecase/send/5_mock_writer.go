// +build dev test

package send

import (
	"errors"
)

type MockWriter struct{}

//
// Mock methods here, with conditionals enabling testing of return values
//
func (m *MockWriter) SubmitSendLinkRequest(email string, linkID string, sessionID string) error {

	if email == "fail@datastore.error" {
		return errors.New("datastore error")
	}

	return nil
}
