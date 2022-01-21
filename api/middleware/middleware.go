package middleware

import (
	"sort"

	"github.com/gin-gonic/gin"
)

var (
	entries []Entry
)

type Middleware interface {
	Init() ([]gin.HandlerFunc, error)
}

type Entry struct {
	m     Middleware
	order int
}

func (e *Entry) Middleware() Middleware {
	return e.m
}

func (e *Entry) Order() int {
	return e.order
}

func AddMiddlewares(es ...Entry) {
	entries = append(entries, es...)
}

func NewEntry(m Middleware, order int) Entry {
	return Entry{
		m:     m,
		order: order,
	}
}

func Entries() []Entry {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].order < entries[j].order
	})
	return entries
}
