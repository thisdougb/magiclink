// +build dev test

package send

import ()

type MockRepository struct {
	MockReader
	MockWriter
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
