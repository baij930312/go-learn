package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/common/util"
)

func serverProcMesLogin(conn net.Conn, msg message.Message) error {
	data := msg.Data
	var loginMes message.LoginMes
	var resMsg message.Message
	var loginRes message.LoginResMes
	resMsg.Type = message.LoginResMesType

	err := json.Unmarshal([]byte(data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return err
	}
	if loginMes.UserId == "100" && loginMes.Password == "123456" {
		loginRes.Code = 200
		data, err := json.Marshal(loginRes)
		if err != nil {
			fmt.Println("json.Marshal", err)
			return err

		}
		resMsg.Data = string(data)

		data, err = json.Marshal(resMsg)
		if err != nil {
			fmt.Println("json.Marshal", err)
			return err

		}
		util.WritePkg(conn, data)
	} else {
		loginRes.Code = 500
		data, err := json.Marshal(loginRes)
		if err != nil {
			fmt.Println("json.Marshal", err)
			return err

		}
		resMsg.Data = string(data)

		data, err = json.Marshal(resMsg)
		if err != nil {
			fmt.Println("json.Marshal", err)
			return err

		}
		util.WritePkg(conn, data)
	}
	return nil
}

func serverProcMes(conn net.Conn, msg message.Message) {
	switch msg.Type {
	case message.LoginMesType:
		serverProcMesLogin(conn, msg)
	case message.LoginResMesType:

	}
}

func run(conn net.Conn) {
	defer conn.Close()
	for {
		msg, err := util.ReadPkg(conn)
		if err != nil {
			fmt.Printf("read pkg  err=%v", err)
			if err == io.EOF {
				fmt.Println("客户端关闭", err)
				break
			}
		}

		serverProcMes(conn, msg)
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
