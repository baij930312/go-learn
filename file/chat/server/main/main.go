package main

import (
	"fmt"
	"net"

	"go_code/learn/file/chat/server/process"
)

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
			process := process.Processer{
				Conn: conn,
			}
			go process.Run()
		}
	}

}
