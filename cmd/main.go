package main

import (
	"github.com/alexkarpovich/quiqstee-user/database"
)

func main() {

  database.InitDB()

  go StartApiServer();
  go StartGrpcServer();

  for {}
}
