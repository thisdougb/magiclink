package send

import (
	"encoding/json"
	"errors"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/entity/sendrequest"
	"log"
)

// Submit a Send task
func (s *Service) Send(email string) error {

	var cfg *config.Config // dynamic config settings

	logins, err := s.repo.GetLoginAttempts(email, cfg.ValueAsInt("RATE_LIMIT_TIME_PERIOD_MINS"))
	if err != nil {
		log.Println("failed getLoginRequests for", email, ":", err.Error())
		if err.Error() == "too many requests" {
			return errors.New("too many requests")
		}
		return errors.New("GetLoginAttempts datastore error")
	}

	if len(logins) >= cfg.ValueAsInt("RATE_LIMIT_MAX_SEND_REQUESTS") {
		log.Println("email is rate limited:", email)
		return errors.New("email address is rate limited")
	}

	sr := sendrequest.NewSendRequest(email)
	if email != sr.Email {
		log.Println("email is invalid:", email)
		return errors.New("email invalid")
	}

	data, err := json.Marshal(sr)
	if err != nil {
		return errors.New("cannot marshal linkrequest: " + err.Error())
	}

	// if stats permit, submit request
	ttlSeconds := 60 * cfg.ValueAsInt("MAGICLINK_EXPIRES_MINS") // seconds * minutes

	err = s.repo.StoreAuthID(sr.Email, sr.MagicLinkID, ttlSeconds)
	if err != nil {
		return err
	}

	err = s.repo.SubmitSendLinkRequest(string(data))
	if err != nil {
		log.Println("failed to submit send link request for", sr.Email, ":", err.Error())
		return err
	}

	log.Println("submitted send link request for", sr.Email)
	return nil
}
