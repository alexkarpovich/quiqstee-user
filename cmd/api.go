package main

import (
	"os"
  "fmt"
	"github.com/alexkarpovich/quiqstee-user/routing"
)

func StartApiServer() {
	address := fmt.Sprintf("%s:%s", os.Getenv("API_HOST"), os.Getenv("API_PORT"))
  fmt.Sprintf("User API server listening %s", address)
  routing.ListenAndServe(address)
}
