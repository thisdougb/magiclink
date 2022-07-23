//go:build dev || test

package poll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// we can't really test the html output here
var TestNotifySendQItems = []struct {
	comment string
	job     string
}{
	{
		comment: "valid request",
		job:     `ts=1618906121,email=me@my.com,magic_link=http://localhost:8080/session/magic-link/authenticate/Jv3pR1eRnAoELJF8J6vZcwBRVzz6APgp8708EUJQHlf1qbwHFWNAkGwul8Y1MiFv`,
	},
}

func TestMagicLinkAuthenticateWeb(t *testing.T) {

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	for _, item := range TestNotifySendQItems {

		job := s.GetNextTask()

		assert.Equal(t, item.job, job)
	}
}
