package owner

import ()

//Repository interface pattern
type Repository interface {
	Reader
	Writer
}

type Reader interface {
	GetSessionOwner(session string) (string, error)
}

type Writer interface {
}
