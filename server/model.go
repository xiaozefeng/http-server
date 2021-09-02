package server

import (
	"github/http-server/context"
)

type Server interface {
	Start(address string) error
	Routable
}

type Handler interface {
	ServeHTTP(c*context.Context)
	Routable
}

type Routable interface {
	Route(method string, pattern string, handler func(c *context.Context))
}
