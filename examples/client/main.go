package main

import (
	"fmt"
	"net"
	"github.com/go75/udpx/pack"
)

func main() {
	remoteAddr := &net.UDPAddr{IP: net.IPv4(127,0,0,1), Port: 9999}
	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println(conn==nil)
	_, err = conn.Write(udpx.ObjtoBytes(udpx.Obj{ID: 0, Payload: []byte("hello")}))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 4096)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	println(addr.String()+":"+string(buf[:n]))
}
