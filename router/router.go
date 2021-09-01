package router

import (
	"fmt"
	"github/http-server/context"
	"github/http-server/model"
	"log"
)

func Hello(c *context.Context){
	c.OK(&model.HTTPResponse{
		Data: "hello world",
	})
}

func SignUp(c *context.Context) {
	var sr SignupReq
	err := c.ReadJson(&sr)
	if err != nil {
		_ = c.BadRequest()
		return
	}
	log.Printf("read request body: %+v", sr)

	var res = model.HTTPResponse{
		Data: sr,
	}
	err = c.OK(&res)
	if err != nil {
		fmt.Fprint(c.W, "write resposne error:", err)
	}
}

type SignupReq struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
