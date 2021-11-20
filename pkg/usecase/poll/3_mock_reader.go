// +build test

package poll

import ()

type MockReader struct{}

func (m *MockReader) GetNextTask() (string, error) {

	var job = "ts=1618906121,email=me@my.com,magic_link=http://localhost:8080/session/magic-link/authenticate/Jv3pR1eRnAoELJF8J6vZcwBRVzz6APgp8708EUJQHlf1qbwHFWNAkGwul8Y1MiFv"

	return job, nil
}
