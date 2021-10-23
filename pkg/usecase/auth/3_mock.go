// +build dev test

package auth

import ()

type MockRepository struct {
	MockReader
	MockWriter
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
