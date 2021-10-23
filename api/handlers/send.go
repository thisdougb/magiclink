package handlers

import (
	"encoding/json"
	"net/http"
)

func (env *Env) Send(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// define this close to its usage
	var input struct {
		Email string
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = env.SendService.Send(input.Email)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}

	http.Error(w, "OK", http.StatusOK)
}
