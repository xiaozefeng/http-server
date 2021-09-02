package server

import (
	"github/http-server/context"
)

type HandlerFunc func(c *context.Context)

type Server interface {
	Start(address string) error
	Routable
}

type Handler interface {
	ServeHTTP(c *context.Context)
	Routable
}

type Routable interface {
	Route(method string, pattern string, handler HandlerFunc)
}
