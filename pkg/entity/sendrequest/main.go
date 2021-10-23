package sendrequest

import (
	"github.com/idthings/alphanum"
	"github.com/idthings/validator"
	"time"
)

const (
	MAGICLINK_ID_LENGTH = 64
	SESSION_ID_LENGTH   = 64
)

type SendRequest struct {
	Email       string
	MagicLinkID string
	SessionID   string
	Timestamp   int64
}

func NewSendRequest(email string) *SendRequest {

	// validate email
	if !validator.IsValidEmail(email) {
		return &SendRequest{}
	}

	magiclinkid := alphanum.New(MAGICLINK_ID_LENGTH)
	sessionid := alphanum.New(SESSION_ID_LENGTH)
	timestamp := time.Now().Unix()

	return &SendRequest{email, magiclinkid, sessionid, timestamp}
}
