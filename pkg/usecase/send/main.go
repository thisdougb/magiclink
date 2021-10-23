package send

import (
	"encoding/json"
	"errors"
	"github.com/thisdougb/magiclink/pkg/entity/linkrequest"
    "github.com/thisdougb/magiclink/config"
)

// EnableThing set the status of a Thing
func (s *Service) Send(email string) error {

	lr := linkrequest.NewLinkRequest(email)
	if email != lr.Email {
		return errors.New("email invalid")
	}

	data, err := json.Marshal(lr)
	if err != nil {
		return errors.New("cannot marshal linkrequest: " + err.Error())
	}

	// if stats permit, submit request
    ttlSeconds := 60 * config.LOGINEXPIRES_MINUTES // seconds * minutes

    err = s.repo.StoreAuthID(lr.Email, lr.MagicLinkID, ttlSeconds)
    if err != nil {
        return err
    }

	err = s.repo.SubmitSendLinkRequest(string(data))
	if err != nil {
		return err
	}

	return nil
}
