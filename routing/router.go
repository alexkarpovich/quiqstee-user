package routing

import (
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/urfave/negroni"
  "github.com/alexkarpovich/quiqstee-user/middlewares"
  "github.com/alexkarpovich/quiqstee-user/routing/common"
  "github.com/alexkarpovich/quiqstee-user/routing/accounts"
)

func router() http.Handler {
    baseRouter := mux.NewRouter().StrictSlash(true)
    common.Router(baseRouter)

    apiRouter := baseRouter.PathPrefix("/api").Subrouter()
    accounts.Router(apiRouter)

    return baseRouter
}

func ListenAndServe(address string) error {
  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

  corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)(router())

  n := negroni.Classic()
  n.Use(negroni.Wrap(middlewares.CurrentUser(corsHandler)))

  server := &http.Server{
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
		Handler: n,
		Addr: address,
	}

	return server.ListenAndServe()
}
