package accounts

import (
    "github.com/gorilla/mux"
)

type AccountHandler struct {}

func Router(r *mux.Router) {
    s := r.PathPrefix("/account").Subrouter()
    h := new(AccountHandler)

    s.HandleFunc("/", h.View).Methods("GET")
}
