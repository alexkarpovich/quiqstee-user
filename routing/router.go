package routing

import (
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/urfave/negroni"
  "github.com/alexkarpovich/quiqstee-user/lib"
  "github.com/alexkarpovich/quiqstee-user/middlewares"
  "github.com/alexkarpovich/quiqstee-user/routing/users"
)

func router() http.Handler {
  apiRouter := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()
  apiRouter.HandleFunc("/healthcheck", healthCheck).Methods("GET")

  users.Router(apiRouter)

  return apiRouter
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    lib.SendJson(w, "Still alive!", http.StatusOK)
}

func ListenAndServe(address string) error {
  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

  corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)(router())

  n := negroni.Classic()
  n.Use(negroni.Wrap(middlewares.User(corsHandler)))

  server := &http.Server{
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
		Handler: n,
		Addr: address,
	}

	return server.ListenAndServe()
}
