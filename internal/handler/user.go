package handler

import (
	"encoding/json"
	"net/http"
)

type createUserInput struct {
	Email, Username string
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var in createUserInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.createUser(r.Context(), in.Email, in.Username)
	if err == service.ErrInvalidEmail || err == service.ErrInvalidUsername {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
