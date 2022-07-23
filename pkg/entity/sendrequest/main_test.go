//go:build dev || test

package sendrequest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidEmails(t *testing.T) {

	emailAddresses := []string{
		"user@domain.com",
		"user@sub.domain.com",
	}

	for _, item := range emailAddresses {
		sr := NewSendRequest(item)
		assert.Equal(t, item, sr.Email, fmt.Sprintf("test valid email %s", item))
		assert.Equal(t, 64, len(sr.MagicLinkID), fmt.Sprintf("test magiclinkid len %s", item))
	}
}

func TestInvalidEmails(t *testing.T) {

	emailAddresses := []string{
		"user",
		"@missing.user",
	}

	for _, item := range emailAddresses {
		sr := NewSendRequest(item)
		assert.Equal(t, "", sr.Email, fmt.Sprintf("test invalid email %s", item))
		assert.Equal(t, 0, len(sr.MagicLinkID), fmt.Sprintf("test magiclinkid is 0 %s", item))
	}
}
