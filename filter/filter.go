package filter

import (
	"github/http-server/context"
	"log"
	"time"
)

type Handler func(c *context.Context)

type Filter func(successor Handler) Handler

type FilterChain struct {
	Filters []Filter
}

func NewFilterChain(filters ...Filter) *FilterChain {
	var fs []Filter
	fs = append(fs, filters...)
	return &FilterChain{
		Filters: fs,
	}
}

var _ Filter = MetricFilterBuilder

func MetricFilterBuilder(successor Handler) Handler {
	return func(c *context.Context) {
		start := time.Now()
		successor(c)
		log.Printf("requst method:%s, request url:%s ,spent %d \n", c.R.Method, c.R.URL, time.Since(start).Microseconds())
	}
}
