package main

import (
	"github.com/alexkarpovich/quiqstee-user/service"
)

func main() {

  service.InitDB()

  go StartApiServer();
  go StartGrpcServer();

  for {}
}
