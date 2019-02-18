package accounts

import (
  "github.com/gorilla/mux"
)

type AccountHandler struct {}

func Router(r *mux.Router) {
  s := r.PathPrefix("/account").Subrouter()
  accountHandler := new(AccountHandler)

  s.HandleFunc("/", accountHandler.View).Methods("GET")
  s.HandleFunc("/signup", accountHandler.Signup).Methods("POST")
  s.HandleFunc("/login", accountHandler.Login).Methods("POST")

}
