package linkrequest

import (
	"github.com/idthings/alphanum"
	"github.com/idthings/validator"
)

const (
	MAGICLINK_ID_LENGTH = 64
	SESSION_ID_LENGTH   = 64
)

type LinkRequest struct {
	Email       string
	MagicLinkID string
	SessionID   string
}

func NewLinkRequest(email string) *LinkRequest {

	// validate email
	if !validator.IsValidEmail(email) {
		return &LinkRequest{}
	}

	magiclinkid := alphanum.New(MAGICLINK_ID_LENGTH)
	sessionid := alphanum.New(SESSION_ID_LENGTH)

	return &LinkRequest{email, magiclinkid, sessionid}
}
