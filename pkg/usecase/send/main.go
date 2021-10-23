package send

import (
	"encoding/json"
	"errors"
	"github.com/thisdougb/magiclink/pkg/entity/sendrequest"
    "github.com/thisdougb/magiclink/config"
)

// EnableThing set the status of a Thing
func (s *Service) Send(email string) error {

	sr := sendrequest.NewSendRequest(email)
	if email != sr.Email {
		return errors.New("email invalid")
	}

	data, err := json.Marshal(sr)
	if err != nil {
		return errors.New("cannot marshal linkrequest: " + err.Error())
	}

	// if stats permit, submit request
    ttlSeconds := 60 * config.MAGICLINK_EXPIRES_MINUTES // seconds * minutes

    err = s.repo.StoreAuthID(sr.Email, sr.MagicLinkID, ttlSeconds)
    if err != nil {
        return err
    }

	err = s.repo.SubmitSendLinkRequest(string(data))
	if err != nil {
		return err
	}

	return nil
}
