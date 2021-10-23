package send

import ()

//Repository interface pattern
type Repository interface {
	Reader
	Writer
}

type Reader interface{}

type Writer interface {
	SubmitSendLinkRequest(data string) error
	StoreAuthID(email string, id string, ttlSeconds int) error
}
