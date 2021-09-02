package v1

import (
	"fmt"
	"github/http-server/context"
	"github/http-server/server"
	"net/http"
	"strings"
)

type HandlerBaseOnTree struct {
	root *node
}

func (h *HandlerBaseOnTree) ServeHTTP(c *context.Context) {
	paths := processURL(c.R.URL.Path)
	cur := h.root

	for _, path := range paths {
		if matchedNode := cur.findMatchChild(path); matchedNode == nil {
			c.W.WriteHeader(http.StatusNotFound)
			fmt.Fprint(c.W, "not found")
			return
		} else {
			cur = matchedNode
		}
	}

	if cur.hadnler == nil {
		c.W.WriteHeader(http.StatusNotFound)
		fmt.Fprint(c.W, "not found")
		return
	}

	// handle request
	cur.hadnler(c)
}

func processURL(path string) []string {
	url := strings.Trim(path, "/")
	paths := strings.Split(url, "/")
	return paths
}

func (h *HandlerBaseOnTree) Route(method string, pattern string, handler server.HandlerFunc) {
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

func (node *node) findMatchChild(path string) *node {
	for _, node := range node.children {
		if node.path == path {
			return node
		}
	}
	return nil
}

type node struct {
	path     string
	children []*node
	hadnler  server.HandlerFunc
}
