package handlers

import (
	"github.com/thisdougb/magiclink/pkg/usecase/auth"
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
	Logger      *log.Logger
	SendService *send.Service
	AuthService *auth.Service
}

func (e *Env) createCookie(sessionName string, sessionID string, expiration time.Time) *http.Cookie {

	cookie := http.Cookie{
		Name:     sessionName,
		Value:    sessionID,
		Path:     "/",
		Expires:  expiration,
		SameSite: http.SameSiteStrictMode,
	}

	return &cookie
}
