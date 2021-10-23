package auth

import ()

//Repository interface pattern
type Repository interface {
	Reader
	Writer
}

type Reader interface {
}

type Writer interface {
	// when authenticating we lookup the magic link id and return the account id,
	// expiring the key at the same time
	GetExpireAccountFromID(magiclinkid string) (string, error)

	StoreSessionID(email string, sessionID string, ttlSeconds int) error
}
