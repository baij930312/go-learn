package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func run(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if length != 4 || err != nil {
			fmt.Printf("读取没有成功 err=%v", err)
			return //!!!
		}

		fmt.Println(binary.BigEndian.Uint32(buf[:length]))
	}
}

func main() {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("error")
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
		} else {
			fmt.Println("远程地址", conn.RemoteAddr().String())
			go run(conn)
		}
	}

}
