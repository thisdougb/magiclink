// +build dev test

package send

import (
	"errors"
	"github.com/thisdougb/magiclink/config"
)

type MockReader struct{}

func (m *MockReader) GetLoginAttempts(email string, sinceMinutes int) ([]string, error) {

	var cfg *config.Config // dynamic config settings

	logins := []string{}

	if email == "getloginsattempts@datastore.error" {
		return logins, errors.New("datastore error")
	}

	if email == "isratelimited@create.session" {
		for i := 0; i <= cfg.RATE_LIMIT_MAX_SEND_REQUESTS()+1; i++ {
			logins = append(logins, "test data")
		}
		return logins, nil
	}

	return logins, nil
}
