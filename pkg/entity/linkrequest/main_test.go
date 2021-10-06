package linkrequest

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
		lr := NewLinkRequest(item)
		assert.Equal(t, item, lr.Email, fmt.Sprintf("test valid email %s", item))
		assert.Equal(t, 64, len(lr.MagicLinkID), fmt.Sprintf("test magiclinkid len %s", item))
		assert.Equal(t, 64, len(lr.SessionID), fmt.Sprintf("test sessionid len %s", item))
	}
}

func TestInvalidEmails(t *testing.T) {

	emailAddresses := []string{
		"user",
		"@missing.user",
	}

	for _, item := range emailAddresses {
		lr := NewLinkRequest(item)
		assert.Equal(t, "", lr.Email, fmt.Sprintf("test invalid email %s", item))
		assert.Equal(t, 0, len(lr.MagicLinkID), fmt.Sprintf("test magiclinkid is 0 %s", item))
		assert.Equal(t, 0, len(lr.SessionID), fmt.Sprintf("test sessionid is 0 %s", item))
	}
}
