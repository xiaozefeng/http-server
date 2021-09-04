package v3

import (
	"errors"
	"fmt"
	"github/http-server/context"
	"github/http-server/server"
	"log"
	"net/http"
	"sort"
	"strings"
)

type nodeType int

const (
	// ROOT node
	ROOT nodeType = iota

	// *
	ANY

	// PATH param node
	PATH

	// regexp
	REG

	// static
	STATIC
)

const anyIdentifier = "*"

type matchFunc func(path string, c *context.Context) bool

type node struct {
	path      string
	children  []*node
	hadnler   server.HandlerFunc
	matchFunc matchFunc
	nodeType  nodeType
}

func newStaticNode(path string) *node {
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *context.Context) bool {
			return path == p && path != anyIdentifier
		},
		nodeType: STATIC,
		path:     path,
	}
}

func newRootNode(method string) *node {
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *context.Context) bool {
			panic("never call me ")
		},
		nodeType: ROOT,
		path:     method,
	}
}

var ErrorInvalidRouterPattern = errors.New("invalid router pattern")
var ErrorInvalidMethod = errors.New("invalid method")

const (
	pathSeparator = "/"
	rootPath      = "/"
)

type HandlerBaseOnTree struct {
	// root *node
	forest map[string]*node
}

var supportedMethods = [4]string{
	http.MethodGet,
	http.MethodPost,
	http.MethodDelete,
	http.MethodPut,
}

func NewHandlerBaseOnTree() server.Handler {
	var forest = make(map[string]*node, len(supportedMethods))
	for _, method := range supportedMethods {
		forest[method] = newRootNode(method)
	}
	return &HandlerBaseOnTree{
		forest: forest,
	}
}

func (h *HandlerBaseOnTree) validatePattern(pattern string) error {
	if pos := strings.Index(pattern, anyIdentifier); pos > 0 {
		if pos != len(pattern)-1 || pattern[pos-1] != '/' {
			return ErrorInvalidRouterPattern
		}
	}
	return nil
}

func (h *HandlerBaseOnTree) ServeHTTP(c *context.Context) {
	// handle request
	if handler := h.findRouter(c.R.Method, c.R.URL.Path, c); handler != nil {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		fmt.Fprint(c.W, "not found\n")
	}
}

func (h *HandlerBaseOnTree) findRouter(method, pattern string, c *context.Context) server.HandlerFunc {
	paths := processURL(pattern)
	var cur, ok = h.forest[method]
	if !ok {
		return nil
	}

	for _, path := range paths {
		if matchedHandler := cur.findMatchChild(path, c); matchedHandler == nil {
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
	log.Printf("register route,method:%s path: %s",method, pattern)

	err := h.validatePattern(pattern)
	if err != nil {
		panic(err)
	}

	paths := processURL(pattern)
	var cur, ok = h.forest[method]
	if !ok {
		panic("invalid method")
	}

	for index, path := range paths {
		if mathedNode := cur.findMatchChild(path, nil); mathedNode != nil {
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

func newParamNode(path string) *node {
	var pathName = path[1:]
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *context.Context) bool {
			if c != nil {
				c.PathParams[pathName] = p
			}
			return p != anyIdentifier
		},
		nodeType: PATH,
	}
}

func newNode(path string) *node {
	if path == anyIdentifier {
		return newAnyNode()
	}
	if strings.HasPrefix(path, ":") {
		return newParamNode(path)
	}
	return newStaticNode(path)
}

func newAnyNode() *node {
	return &node{
		matchFunc: func(path string, c *context.Context) bool {
			return true
		},
		nodeType: ANY,
		path:     anyIdentifier,
	}
}

func (nd *node) findMatchChild(path string, c *context.Context) *node {
	var candidates = make([]*node, 0, 2)
	for _, child := range nd.children {
		if child.matchFunc(path, c) {
			candidates = append(candidates, child)
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].nodeType < candidates[j].nodeType
	})
	return candidates[len(candidates)-1]
}
