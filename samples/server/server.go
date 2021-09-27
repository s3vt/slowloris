package main

import (
	"log"

	"github.com/sapvs/slowloris/samples"
)

func main() {
	bas := &samples.BasicAuthServer{
		BaseServer: &samples.BaseServer{
			Address: samples.Address},
		Username: samples.BasicAuthUser,
		Password: samples.BasicAuthPassword}
	ds := &samples.DefaultServer{BaseServer: &samples.BaseServer{Address: "localhost:8080"}}

	log.Fatalf("Server done with %v", startServer(bas))
	log.Fatalf("Server done with %v", startServer(ds))

}

func startServer(server samples.Server) error {
	log.Printf("Starting %T server ", server)
	return server.Start()
}
