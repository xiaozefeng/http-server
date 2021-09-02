package sdk

import (
	"github/http-server/context"
	"github/http-server/filter"
	"github/http-server/sdk/handlermapping"
	"github/http-server/server"
	"net/http"
)

type Server struct {
	Name    string
	Handler server.Handler
	Root    filter.Handler
}

func (s *Server) Route(method, pattern string, hadnler server.HandlerFunc) {
	s.Handler.Route(method, pattern, hadnler)
}

func (s *Server) Start(address string) error {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var c = context.NewContext(rw, r)
		s.Root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHTTPServer(name string, fs ...filter.Filter) server.Server {
	var handler = handlermapping.New()
	var root = handler.ServeHTTP
	for i := len(fs)-1; i >= 0; i-- {
		f := fs[i]
		root = f(root)
	}

	return &Server{
		Name:    name,
		Handler: handler,
		Root:    root,
	}
}
