package sendrequest

import (
	"github.com/idthings/alphanum"
	"github.com/idthings/validator"
	"github.com/thisdougb/magiclink/config"
	"time"
)

type SendRequest struct {
	Email       string
	MagicLinkID string
	Timestamp   int64
}

func NewSendRequest(email string) *SendRequest {

	var cfg *config.Config // dynamic config settings

	// validate email
	if !validator.IsValidEmail(email) {
		return &SendRequest{}
	}

	magiclinkid := alphanum.New(cfg.MAGICLINK_LENGTH())
	timestamp := time.Now().Unix()

	return &SendRequest{email, magiclinkid, timestamp}
}
