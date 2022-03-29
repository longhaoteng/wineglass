package middleware

import (
	"sort"

	"github.com/gin-gonic/gin"
)

var entries []Entry

type Middleware interface {
	Init() ([]gin.HandlerFunc, error)
}

type Entry struct {
	mw    Middleware
	order int
}

func (e *Entry) Middleware() Middleware {
	return e.mw
}

func (e *Entry) Order() int {
	return e.order
}

func AddMiddlewares(es ...Entry) {
	entries = append(entries, es...)
}

func NewEntry(m Middleware, orders ...int) Entry {
	order := 0
	if len(orders) > 0 {
		order = orders[0]
	}

	return Entry{
		mw:    m,
		order: order,
	}
}

func Entries() []Entry {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].order < entries[j].order
	})
	return entries
}
