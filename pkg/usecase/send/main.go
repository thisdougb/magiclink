package send

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/entity/sendrequest"
)

// Submut a Send task
func (s *Service) Send(email string) error {

	var cfg *config.Config // dynamic config settings

	// get the logins for the rate limiting period of time
	logins, err := s.repo.GetLoginAttempts(email, cfg.RATE_LIMIT_TIME_PERIOD_MINS())
	if err != nil {
		return errors.New("GetLoginAttempts datastore error")
	}

	fmt.Println("logsin:", len(logins))
	if len(logins) >= cfg.RATE_LIMIT_MAX_SEND_REQUESTS() {
		return errors.New("email address is rate limited")
	}

	err = s.repo.LogLoginAttempt(email)
	if err != nil {
		return errors.New("LogLoginAttempt datastore error")
	}

	sr := sendrequest.NewSendRequest(email)
	if email != sr.Email {
		return errors.New("email invalid")
	}

	data, err := json.Marshal(sr)
	if err != nil {
		return errors.New("cannot marshal linkrequest: " + err.Error())
	}

	// if stats permit, submit request
	ttlSeconds := 60 * cfg.MAGICLINK_EXPIRES_MINS() // seconds * minutes

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
