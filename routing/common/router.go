package common

import (
    "net/http"
    "github.com/gorilla/mux"
)

type CommonHandler struct {}

func Router(r *mux.Router) {
    commonHandler := new(CommonHandler)

    r.NotFoundHandler = http.HandlerFunc(commonHandler.Handler404)
    r.HandleFunc("/healthcheck", commonHandler.HealthCheck).Methods("GET")
}
