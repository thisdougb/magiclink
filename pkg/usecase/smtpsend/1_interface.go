package smtpsend

import ()

//Repository interface
type Repository interface {
	Reader
	Writer
}

//Reader datastore interface
type Reader interface{}

//Writer user writer
type Writer interface{}
