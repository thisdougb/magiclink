package send

import ()

//Repository interface pattern
type Repository interface {
	Reader
	Writer
}

type Reader interface{}

type Writer interface {
	SubmitSendLinkRequest(email string, linkID string, sessionID string) error
}
