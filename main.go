package main

import (
	"github/http-server/filter"
	"github/http-server/router"
	"github/http-server/sdk"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	var filters = []filter.Filter{filter.MetricFilter}
	s := sdk.NewHTTPServer("http-server", filters...)

	s.Route(http.MethodPost, "/signup", router.SignUp)
	s.Route(http.MethodGet, "/signup/hello/*", router.Hello)
	s.Route(http.MethodGet, "/hello", router.Hello)

	log.Print("start http server on address 8080\n")

	go func() {
		if err := s.Start(":8080"); err != nil {
			log.Fatalf("start server error: %v", err)
		}
	}()

	WaitForShutdown()
}

func WaitForShutdown() {
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan)

	for {
		select {
		case sig := <-signChan:
			log.Printf("get signal: %s\n", sig)
			time.AfterFunc(time.Minute, func() {
				log.Print("timeout 1 minute, applicaiton shutdown gracefully\n")
				os.Exit(1)
			})
			log.Print("applicaiton shutdown gracefully\n")
			os.Exit(0)
		}

	}
}
