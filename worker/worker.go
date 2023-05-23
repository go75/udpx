package worker

import (
	"github.com/go75/udpx/dispatcher"
	"github.com/go75/udpx/router"
)

func InitWorkerPool(workerSize uint32) {	
	for workerSize > 0 {
		go func() {
			for req := range dispatcher.Dispatcher {
				router.Handle(req)
			}
		}()
		workerSize--
	}
}