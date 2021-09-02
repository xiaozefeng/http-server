package v1

import (
	"github/http-server/context"
	"github/http-server/server"
)

type HandlerBaseOnTree struct{
	root *node
}


func (h *HandlerBaseOnTree) ServeHTTP(c *context.Context) {

}

func (h *HandlerBaseOnTree) Route(method string, pattern string, handler func(c *context.Context)) {

}


type node struct{
	path string
	children []*node

	hadnler server.HandlerFunc
}
