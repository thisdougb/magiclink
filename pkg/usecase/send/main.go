package send

import (
	"errors"
	"github.com/thisdougb/magiclink/pkg/entity/linkrequest"
)

// EnableThing set the status of a Thing
func (s *Service) Send(email string) error {

	lr := linkrequest.NewLinkRequest(email)
	if email != lr.Email {
		return errors.New("email invalid")
	}

	// read email stats

	// if stats permit, submit request
	err := s.repo.SubmitSendLinkRequest(lr.Email, lr.MagicLinkID, lr.SessionID)
	if err != nil {
		return err
	}

	return nil
}
