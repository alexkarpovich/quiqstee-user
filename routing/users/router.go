package users

import (
  "github.com/gorilla/mux"
)

type UserHandler struct {}

func Router(r *mux.Router) {
  s := r.PathPrefix("/users").Subrouter()
  userHandler := new(UserHandler)

  s.HandleFunc("/signup", userHandler.Signup).Methods("POST")
  s.HandleFunc("/login", userHandler.Login).Methods("POST")
  s.HandleFunc("/view", userHandler.View).Methods("GET")
}
