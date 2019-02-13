package routing

import (
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/alexkarpovich/quiqstee-user/routing/users"
)

func router() http.Handler {
  apiRouter := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()
  users.Router(apiRouter)
  return apiRouter
}

func ListenAndServe(address string) error {
  server := &http.Server{
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
		Handler: router(),
		Addr: address,
	}

	return server.ListenAndServe()
}
