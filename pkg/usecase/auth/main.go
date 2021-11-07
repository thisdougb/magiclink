package auth

import (
	"errors"
	"github.com/idthings/alphanum"
	"github.com/thisdougb/magiclink/config"
)

// Auth authenticates an inbound magic link url
func (s *Service) Auth(magiclinkid string) (string, error) {

	var cfg *config.Config // dynamic config settings

	if !alphanum.IsValidAlphaNum(magiclinkid, cfg.MAGICLINK_LENGTH()) {
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

	sessionID := alphanum.New(cfg.SESSION_ID_LENGTH())
	ttlSeconds := 60 * cfg.SESSION_EXPIRES_MINS() // seconds * minutes

	err = s.repo.StoreSessionID(email, sessionID, ttlSeconds)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}
