package handlermapping

import (
	"fmt"
	"github/http-server/context"
	"github/http-server/server"
	"log"
	"net/http"
)

type HandlerMappingOnMap struct {
	Hadnlers map[string]func(c *context.Context)
}

func New() server.Handler {
	return &HandlerMappingOnMap{
		Hadnlers: make(map[string]func(c *context.Context)),
	}
}

func (h *HandlerMappingOnMap) Route(method, pattern string, handler func(c *context.Context)) {
	var key = h.Key(method, pattern)
	log.Printf("register route, method:%s, path:%s", method, pattern)
	h.Hadnlers[key] = handler
}

func (h *HandlerMappingOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var key = h.Key(r.Method, r.URL.Path)
	if handler, ok := h.Hadnlers[key]; ok {
		c := context.NewContext(w, r)
		handler(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, "not any router math")
	}
}

func (h *HandlerMappingOnMap) Key(method, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
