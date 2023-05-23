package router

import "github.com/go75/udpx/mod"

var m = make(map[uint32]func(mod.Request), 0)

func Put(id uint32, fn func(mod.Request)) {
	m[id] = fn
}

func Handle(req mod.Request) {
	m[req.ID](req)
}