package sendrequest

import (
	"github.com/idthings/alphanum"
	"github.com/idthings/validator"
	"github.com/thisdougb/magiclink/config"
	"time"
)

const (
	SESSION_ID_LENGTH = 64
)

type SendRequest struct {
	Email       string
	MagicLinkID string
	Timestamp   int64
}

func NewSendRequest(email string) *SendRequest {

	// validate email
	if !validator.IsValidEmail(email) {
		return &SendRequest{}
	}

	magiclinkid := alphanum.New(config.MAGICLINK_ID_LENGTH)
	timestamp := time.Now().Unix()

	return &SendRequest{email, magiclinkid, timestamp}
}
