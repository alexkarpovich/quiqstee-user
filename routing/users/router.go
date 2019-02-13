package routing

import (
  "github.com/gorilla/mux"
)

type UserHandler struct {}

func UserRouter(r *mux.Router) {
  s := r.PathPrefix("/users").Subrouter()
  userHandler := new(UserHandler)

  s.HandleFunc("/signup", userHandler.Signup).Methods("POST")
  s.HandleFunc("/login", userHandler.Login).Methods("POST")
  s.HandleFunc("/logout", userHandler.Login).Methods("POST")
}

func Dno() {}
