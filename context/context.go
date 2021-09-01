package context

import (
	"encoding/json"
	"fmt"
	"github/http-server/model"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}

func (c *Context) ReadJson(obj interface{}) error {
	return json.NewDecoder(c.R.Body).Decode(obj)
}

func (c *Context) WriteResponse(code int, res *model.HTTPResponse) error {
	c.W.WriteHeader(code)
	if res != nil {
		content, err := json.Marshal(&res)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(c.W, string(content))
		return err
	}
	return nil
}

func (c *Context) OK(res *model.HTTPResponse) error {
	res.Code = 200
	res.Msg = "success"
	return c.WriteResponse(http.StatusOK, res)
}

func (c *Context) BadRequest() error {
	return c.WriteResponse(http.StatusBadRequest, nil)
}
