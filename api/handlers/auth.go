package handlers

import (
	"github.com/thisdougb/magiclink/config"
	"net/http"
	"strings"
	"time"
)

func (env *Env) Auth(w http.ResponseWriter, r *http.Request) {

	var cfg *config.Config // dynamic config settings

	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	magiclinkid := strings.TrimPrefix(r.URL.Path, "/auth/")

	sessionID, err := env.AuthService.Auth(magiclinkid)
	if err != nil {
		// we don't set a cookie here
		if err.Error() == "magic link not found" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		if err.Error() == "invalid magic link id" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	expiresAtTime := time.Now().Add(time.Duration(cfg.SESSION_EXPIRES_MINS()) * time.Minute)
	cookie := env.createCookie(cfg.SESSION_NAME(), sessionID, expiresAtTime)

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
