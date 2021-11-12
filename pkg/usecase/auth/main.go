package auth

import (
	"errors"
	"github.com/idthings/alphanum"
	"github.com/thisdougb/magiclink/config"
	"log"
)

// Auth authenticates an inbound magic link url
func (s *Service) Auth(magiclinkid string) (string, error) {

	var cfg *config.Config // dynamic config settings

	if !alphanum.IsValidAlphaNum(magiclinkid, cfg.ValueAsInt("MAGICLINK_LENGTH")) {
		log.Println("magiclink id is invalid:", magiclinkid)
		return "", errors.New("invalid magic link id")
	}

	email, err := s.repo.GetExpireAccountFromID(magiclinkid)
	if err != nil {
		return "", err
	}

	// possible fake link, expired link, etc.
	if len(email) == 0 {
		log.Println("magiclink id not found:", magiclinkid)
		return "", errors.New("magic link not found")
	}

	sessionID := alphanum.New(cfg.ValueAsInt("SESSION_ID_LENGTH"))
	ttlSeconds := 60 * cfg.ValueAsInt("SESSION_EXPIRES_MINS") // seconds * minutes

	err = s.repo.StoreSessionID(email, sessionID, ttlSeconds)
	if err != nil {
		log.Println("failed to store session id for", email, ":", err.Error())
		return "", err
	}

	return sessionID, nil
}
