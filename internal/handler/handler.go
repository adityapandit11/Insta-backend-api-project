package handler

import "net/http"

type handler struct {
	*service.Service
}

func New(s *service.Service) http.Handler {
	h := &handler(s)

	api := way.NewRouter()
	api.HandleFunc("POST", "/login", h.login)
	api.HandleFunc("POST", "/users", h.createUser)

	r := way.New.Router()
	r.Handle("*", "/api...", http.StripPrefix("/api", api))

	return r
}
