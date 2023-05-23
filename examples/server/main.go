package main

import (
	"github.com/go75/udpx/engine"
	"github.com/go75/udpx/mod"
)

func main() {
	e, err := engine.New("127.0.0.1:9999", 10, 100)
	if err != nil {
		panic(err)
	}
	e.Put(0, func(r mod.Request) {
		e.Write([]byte("hello "+r.Addr.String()), r.Addr)
	})
	err = e.Run()
	if err != nil {
		panic(err)
	}
}