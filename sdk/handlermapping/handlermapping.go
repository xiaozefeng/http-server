package handlermapping

import (
	"fmt"
	"github/http-server/context"
	"github/http-server/server"
	"log"
	"net/http"
	"sync"
)

type HandlerMappingOnMap struct {
	Hadnlers sync.Map
}

var _ server.Handler = &HandlerMappingOnMap{}

func New() server.Handler {
	return &HandlerMappingOnMap{
		Hadnlers: sync.Map{},
	}
}

func (h *HandlerMappingOnMap) Route(method, pattern string, handler server.HandlerFunc) {
	var key = h.Key(method, pattern)
	log.Printf("register route, method:%s, path:%s", method, pattern)
	// h.Hadnlers[key] = handler
	h.Hadnlers.Store(key, handler)
}

func (h *HandlerMappingOnMap) ServeHTTP(c *context.Context) {
	var key = h.Key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.Hadnlers.Load(key); ok {
		c := context.NewContext(c.W, c.R)
		handler.(server.HandlerFunc)(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(c.W, "not any router math")
	}
}

func (h *HandlerMappingOnMap) Key(method, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
