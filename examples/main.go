package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	for i:=10000;i<10010;i++ {
		go func(j int){
			addr, err := net.ResolveUDPAddr("udp", "10.105.12.220:" + strconv.Itoa(j))
			if err != nil {
				panic(err)
			}
			conn, err := net.ListenUDP("udp", addr)
			if err != nil{
				panic(err)
			}
			buf := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}  
			println("read from "+addr.String()+" : "+string(buf[:n]))
		}(i)
	}
	broadcast()
}


func broadcast()  {

	// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
	laddr := net.UDPAddr{
		IP:   net.IPv4(10, 105, 12, 220),
		Port: 3000,
	}

	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 3000,
	}

	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}

	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil{
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
