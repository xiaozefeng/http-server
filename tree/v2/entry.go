package v2

import (
	"errors"
	"fmt"
	"github/http-server/context"
	"github/http-server/server"
	"log"
	"net/http"
	"strings"
)

var ErrorInvalidRouterPattern = errors.New("invalid router pattern")

const (
	pathSeparator = "/"
	rootPath      = "/"
)

type HandlerBaseOnTree struct {
	root *node
}

func NewHandlerBaseOnTree() server.Handler {
	return &HandlerBaseOnTree{
		root: &node{path: rootPath},
	}
}

func (h *HandlerBaseOnTree) validatePattern(pattern string) error {
	if pos := strings.Index(pattern, "*"); pos > 0 {
		if pos != len(pattern)-1 || pattern[pos-1] != '/' {
			return ErrorInvalidRouterPattern
		}
	}
	return nil
}

func (h *HandlerBaseOnTree) ServeHTTP(c *context.Context) {
	// handle request
	if matchedHandler := h.findRouter(c.R.URL.Path); matchedHandler != nil {
		matchedHandler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		fmt.Fprint(c.W, "not found\n")
	}
}

func (h *HandlerBaseOnTree) findRouter(urlPath string) server.HandlerFunc {
	paths := processURL(urlPath)
	cur := h.root

	for _, path := range paths {
		if matchedHandler := cur.findMatchChild(path); matchedHandler == nil {
			// not found
			return nil
		} else {
			cur = matchedHandler
		}
	}
	return cur.hadnler
}

func processURL(path string) []string {
	url := strings.Trim(path, pathSeparator)
	paths := strings.Split(url, pathSeparator)
	return paths
}

func (h *HandlerBaseOnTree) Route(method string, pattern string, handler server.HandlerFunc) {
	log.Printf("register route, path: %s", pattern)
	err := h.validatePattern(pattern)
	if err != nil {
		panic(err)
	}

	paths := processURL(pattern)
	cur := h.root

	for index, path := range paths {
		if mathedNode := cur.findMatchChild(path); mathedNode != nil {
			cur = mathedNode
		} else {
			h.createSubTree(cur, paths[index:], handler)
			break
		}
	}
	cur.hadnler = handler
}

func (h *HandlerBaseOnTree) createSubTree(root *node, paths []string, handler server.HandlerFunc) {
	cur := root
	for _, path := range paths {
		node := newNode(path)
		cur.children = append(cur.children, node)
		cur = node
	}
	cur.hadnler = handler
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

func (nd *node) findMatchChild(path string) *node {
	var wildcardNode *node
	for _, n := range nd.children {
		if n.path == path && n.path != "*" {
			return n
		}

		if n.path == "*" {
			wildcardNode = n
		}
	}
	return wildcardNode
}

type node struct {
	path     string
	children []*node
	hadnler  server.HandlerFunc
}
