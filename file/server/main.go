package main

import (
	"fmt"
	"net"
)

func run(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		lenght, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		fmt.Println(string(buf[:lenght]))
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
