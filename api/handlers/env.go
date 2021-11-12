package handlers

import (
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/usecase/auth"
	"github.com/thisdougb/magiclink/pkg/usecase/owner"
	"github.com/thisdougb/magiclink/pkg/usecase/send"
	"log"
	"net/http"
	"time"
)

/*
   The Env struct allows us to pass the datastore pointer around,
   it also allows us to inject mocks in usecase packages.
*/

type Env struct {
	Logger       *log.Logger
	SendService  *send.Service
	AuthService  *auth.Service
	OwnerService *owner.Service
}

func (e *Env) createCookie(sessionName string, sessionID string, expiresAtTime time.Time) *http.Cookie {

	cookie := http.Cookie{
		Name:     sessionName,
		Value:    sessionID,
		Path:     "/",
		Expires:  expiresAtTime,
		SameSite: http.SameSiteStrictMode,
	}

	return &cookie
}

func (e *Env) GetURLPrefix() string {

	var cfg *config.Config // dynamic config settings

	urlPrefix := cfg.ValueAsStr("URL_PREFIX")

	// strip trailing / from url prefix if it exists
	urlPrefixLength := len(urlPrefix)
	if urlPrefixLength > 0 && urlPrefix[urlPrefixLength-1] == '/' {
		urlPrefix = urlPrefix[:urlPrefixLength-1]
	}
	return urlPrefix
}
