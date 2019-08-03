package main

import (
	"fmt"
	"net"
	"time"

	"go_code/learn/file/chat/server/model"
	"go_code/learn/file/chat/server/process"
)

func initUserDao() {
	model.UserDaoInstance = model.NewUserDao(pool)
}

func init() {
	initPoll("localhost:6379", 16, 0, time.Second*300)
	initUserDao()
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
			process := process.Processer{
				Conn: conn,
			}
			go process.Run()
		}
	}

}
