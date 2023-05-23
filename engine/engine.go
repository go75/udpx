package engine

import (
	"errors"
	"fmt"
	"net"
	"github.com/go75/udpx/dispatcher"
	"github.com/go75/udpx/mod"
	"github.com/go75/udpx/router"
	"github.com/go75/udpx/utils"
	"github.com/go75/udpx/worker"
)

type Engine struct {
	conn *net.UDPConn
	writerSize uint32
	isStart bool
	ErrChan chan error
}

func (e *Engine)Write(b []byte, addr *net.UDPAddr) (int, error) {
	return e.conn.WriteToUDP(b, addr)
}

func New(address string, writerSize, dispatcherSize uint32) (*Engine, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	dispatcher.InitDispatcher(dispatcherSize)
	return &Engine{
		conn: conn,
		writerSize: writerSize,
		ErrChan: make(chan error),
	}, nil
}

func (e *Engine) Addr() net.Addr {
	return e.conn.LocalAddr()
}

func (e *Engine) IsStart() bool {
	return e.isStart
}

func (e *Engine) Put(id uint32, fn func(mod.Request)) {
	if e.isStart {
		e.ErrChan <- errors.New("server is starting, can not regist handler")
	}
	router.Put(id, fn)
}

func (e *Engine) Run() error {
	e.isStart = true
	worker.InitWorkerPool(e.writerSize)
	fmt.Println("listening on " + e.conn.LocalAddr().String())
	go func() {
		for {
			buf := make([]byte, 520)
			n, addr, err := e.conn.ReadFromUDP(buf)
			if err != nil {
				return
			}
			
			dispatcher.Dispatcher <- mod.Request {
				Addr: addr,
				Obj: mod.Obj {
					ID: utils.BytesToUint32(buf[:4]),
					Payload: buf[4:n],
				},
			}
		}
	}()
	return <-e.ErrChan
}