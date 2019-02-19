package main

import (
    "sync"
	"github.com/alexkarpovich/quiqstee-user/database"
)

func main() {
    var wg sync.WaitGroup

    database.InitDB()
    wg.Add(1)
    go StartApiServer();
    wg.Add(2)
    go StartGrpcServer();

    wg.Wait()
}
