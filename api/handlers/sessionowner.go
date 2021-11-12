package handlers

import (
	"encoding/json"
	"github.com/thisdougb/magiclink/config"
	"net/http"
	"strings"
)

func (env *Env) Owner(w http.ResponseWriter, r *http.Request) {

	var cfg *config.Config // dynamic config settings

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// define this close to its usage
	var input struct {
		Token   string
		Session string
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// get access tokens currently allowed
	accessTokens := strings.Split(cfg.ValueAsStr("SESSION_OWNER_ACCESS_TOKENS"), ",")
	accessTokenValid := false
	for _, token := range accessTokens {
		if input.Token == token {
			accessTokenValid = true
			break
		}
	}
	if !accessTokenValid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	owner, err := env.OwnerService.SessionOwner(input.Session)
	if err != nil || owner == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var output struct {
		Owner string
	}

	output.Owner = owner

	outputBytes, err := json.Marshal(output)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(outputBytes)
}
