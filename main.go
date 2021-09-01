package main

import (
	"github/http-server/router"
	"github/http-server/sdk"
	"log"
)

func main() {
	s := sdk.NewHTTPServer()
	s.Route("POST", "/signup", router.SignUp)
	s.Route("GET", "/hello", router.Hello)
	log.Print("start http server on address 8080\n")
	if err := s.Start(":8080"); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}
