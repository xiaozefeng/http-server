package sdk

import (
	"net/http"

	"github/http-server/context"
	"github/http-server/sdk/handlermapping"
	"github/http-server/server"
)

type Server struct {
	Name    string
	Handler server.Handler
}

func (s *Server) Route(method, pattern string, hadnler func(c *context.Context)) {
	s.Handler.Route(method, pattern, hadnler)
}

func (s *Server) Start(address string) error {
	return http.ListenAndServe(address, s.Handler)
}

func NewHTTPServer() server.Server {
	return &Server{Name: "sdk server", Handler: handlermapping.New()}
}
