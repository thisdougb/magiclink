// +build dev test

package send

import (
	"errors"
	"github.com/thisdougb/magiclink/config"
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

func (m *MockWriter) GetLoginAttempts(email string, ttlMinutes int) ([]string, error) {

	var cfg *config.Config // dynamic config settings

	logins := []string{}

	if email == "getloginsattempts@datastore.error" {
		return logins, errors.New("datastore error")
	}

	if email == "getloginattempts@toomany.requests" {
		return logins, errors.New("too many requests")
	}
	if email == "isratelimited@create.session" {
		for i := 0; i <= cfg.RATE_LIMIT_MAX_SEND_REQUESTS()+1; i++ {
			logins = append(logins, "test data")
		}
		return logins, nil
	}

	if email == "logloginattempt@datastore.error" {
		return []string{}, errors.New("datastore error")
	}
	return []string{}, nil
}
