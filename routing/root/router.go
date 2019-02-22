package root

import (
    "net/http"
    "github.com/gorilla/mux"
)

type RootHandler struct {}

func Router(r *mux.Router) {
    h := new(RootHandler)

    r.NotFoundHandler = http.HandlerFunc(h.Handler404)
    r.HandleFunc("/healthcheck", h.HealthCheck).Methods("GET")

    r.HandleFunc("/signup", h.Signup).Methods("POST")
    r.HandleFunc("/signup/{token}", h.ConfirmSignup).Methods("POST")
    r.HandleFunc("/login", h.Login).Methods("POST")
}
