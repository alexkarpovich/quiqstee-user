package routing

import (
  "os"
  "time"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/alexkarpovich/quiqstee-user/routing/users"
)

func router() http.Handler {
  apiRouter := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()
  apiRouter.HandleFunc("/healthcheck", healthCheck).Methods("GET")

  users.Router(apiRouter)

  return apiRouter
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Still alive!")
}

func ListenAndServe(address string) error {
  headersOk := handlers.AllowedHeaders([]string{"Authorization"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

  corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)
  loggedRouter := handlers.LoggingHandler(os.Stdout, router())

  server := &http.Server{
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
		Handler: corsHandler(loggedRouter),
		Addr: address,
	}

	return server.ListenAndServe()
}
