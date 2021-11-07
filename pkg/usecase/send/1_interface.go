package send

import ()

//Repository interface pattern
type Repository interface {
	Reader
	Writer
}

type Reader interface {
	GetLoginAttempts(email string, sinceMinutes int) ([]string, error)
}

type Writer interface {
	SubmitSendLinkRequest(data string) error
	StoreAuthID(email string, id string, ttlSeconds int) error

	GetLoginAttempts(email string, sinceMinutes int) ([]string, error)
}
