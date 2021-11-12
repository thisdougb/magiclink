package owner

import (
	"errors"
	"github.com/idthings/alphanum"
	"github.com/thisdougb/magiclink/config"
	"log"
)

// Submit a Send task
func (s *Service) SessionOwner(session string) (string, error) {

	var cfg *config.Config // dynamic config settings

	if !alphanum.IsValidAlphaNum(session, cfg.ValueAsInt("SESSION_ID_LENGTH")) {
		log.Println("session id is invalid:", session)
		return "", errors.New("invalid session id")
	}

	// lookup session id
	owner, err := s.repo.GetSessionOwner(session)
	if err != nil {
		return "", err
	}

	return owner, nil
}
