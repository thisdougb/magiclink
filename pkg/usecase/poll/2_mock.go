// +build dev

package poll

import ()

type MockRepository struct {
	MockReader
	MockWriter
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
