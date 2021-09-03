package main

import (
	"github/http-server/filter"
	"github/http-server/router"
	"github/http-server/sdk"
	"log"
	"net/http"
)

func main() {

	var filters = []filter.Filter{ filter.MetricFilter}
	s := sdk.NewHTTPServer("http-server", filters...)

	s.Route(http.MethodPost, "/signup", router.SignUp)
	s.Route(http.MethodGet, "/signup/hello/*", router.Hello)
	s.Route(http.MethodGet, "/hello", router.Hello)

	log.Print("start http server on address 8080\n")
	if err := s.Start(":8080"); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}
