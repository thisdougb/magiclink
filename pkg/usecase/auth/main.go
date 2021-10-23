package auth

import (
	"errors"
	"github.com/idthings/alphanum"
	"github.com/thisdougb/magiclink/config"
)

// EnableThing set the status of a Thing
func (s *Service) Auth(magiclinkid string) (string, error) {

	if !alphanum.IsValidAlphaNum(magiclinkid, config.MAGICLINK_ID_LENGTH) {
		return "", errors.New("invalid magic link id")
	}

	email, err := s.repo.GetExpireAccountFromID(magiclinkid)
	if err != nil {
		return "", err
	}

	// possible fake link, expired link, etc.
	if len(email) == 0 {
		return "", errors.New("magic link not found")
	}

	sessionID := alphanum.New(config.SESSION_ID_LENGTH)
	ttlSeconds := 60 * config.SESSION_ID_EXPIRES_MINUTES // seconds * minutes

	err = s.repo.StoreSessionID(email, sessionID, ttlSeconds)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}
