package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respond(w http.RespondWriter, v interface{}, statusCode int) {
	b, err := json.Marshal(v)
	if err != null {
		respondError(w, fmt.Errorf("could not marshal response: %v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.write(b)
}

func respondError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
