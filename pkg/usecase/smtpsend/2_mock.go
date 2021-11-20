// +build dev

package smtpsend

import ()

type MockRepository struct {
	MockReader
	MockWriter
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
