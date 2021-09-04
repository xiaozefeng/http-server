package sdk

import (
	"github/http-server/context"
	"github/http-server/filter"
	"github/http-server/server"
	v3 "github/http-server/tree/v3"
	"net/http"
	"sync"
)

type Server struct {
	name    string
	handler server.Handler
	root    filter.Handler
	pool    sync.Pool
}

func (s *Server) Route(method, pattern string, hadnler server.HandlerFunc) {
	s.handler.Route(method, pattern, hadnler)
}

func (s *Server) Start(address string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := s.pool.Get().(*context.Context)
		defer s.pool.Put(c)
		c.Reset(w, r)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHTTPServer(name string, fs ...filter.Filter) server.Server {
	// var handler = handlermapping.New()
	var handler = v3.NewHandlerBaseOnTree()
	var root = handler.ServeHTTP
	for i := len(fs) - 1; i >= 0; i-- {
		f := fs[i]
		root = f(root)
	}

	return &Server{
		name:    name,
		handler: handler,
		root:    root,
		pool: sync.Pool{
			New: func() interface{} {
				return context.NewEmptyContext()
			},
		},
	}
}
