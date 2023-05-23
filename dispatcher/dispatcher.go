package dispatcher

import "github.com/go75/udpx/mod"

var Dispatcher chan mod.Request

func InitDispatcher(dispatcherSize uint32) {
	Dispatcher = make(chan mod.Request, dispatcherSize)
}