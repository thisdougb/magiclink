package poll

import ()

//Repository interface
type Repository interface {
	Reader
	Writer
}

//Reader datastore interface
type Reader interface {
	GetNextTask() (string, error)
}

//Writer user writer
type Writer interface {
}
