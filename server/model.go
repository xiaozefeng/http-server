package server

import (
	"github/http-server/context"
	"net/http"
)

type Server interface {
	Start(address string) error
	Routable
}

type Handler interface {
	http.Handler
	Routable
}

type Routable interface {
	Route(method string, pattern string, handler func(c *context.Context))
}
