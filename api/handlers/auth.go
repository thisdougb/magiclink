package handlers

import (
	"github.com/thisdougb/magiclink/config"
	"net/http"
	"strings"
	"time"
)

func (env *Env) Auth(w http.ResponseWriter, r *http.Request) {

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

	expiration := time.Now().Add(config.HttpSessionTTL)
	cookie := env.createCookie(config.HttpSessionName, sessionID, expiration)

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
