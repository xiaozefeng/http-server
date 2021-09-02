package filter

// import (
// 	"github/http-server/context"
// 	"log"
// )

// type Handler interface {
// 	Process(c *context.Context) bool
// }

// type FilterChain struct {
// 	Handlers []Handler
// }

// func NewFilterChain() FilterChain {
// 	return FilterChain{
// 		Handlers: make([]Handler, 0),
// 	}
// }

// func (fc *FilterChain) AddHandler(h Handler) {
// 	fc.Handlers = append(fc.Handlers, h)
// }

// func (fc *FilterChain) ProcessRequest(c *context.Context) {
// 	for _, handler := range fc.Handlers {
// 		if !handler.Process(c) {
// 			return
// 		}
// 	}
// }

// type ManagerHandler struct {
// }

// func (mh *ManagerHandler) Process(c *context.Context) bool {
// 	log.Print("Manager Hadnler\n")
// 	return true
// }

// type DirectorHandler struct {
// }

// func (mh *DirectorHandler) Process(c *context.Context) bool {
// 	log.Print("Director Handler\n")
// 	return true
// }

// type CEOHandler struct {
// }

// func (mh *CEOHandler) Process(c *context.Context) bool {
// 	log.Print("CEO Handler\n")
// 	return true
// }

// func NewMangerHandler() Handler {
// 	return &ManagerHandler{}
// }

// func NewDirectorHandler() Handler {
// 	return &DirectorHandler{}
// }
// func NewCEOHandler() Handler {
// 	return &CEOHandler{}
// }
