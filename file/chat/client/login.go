package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/common/util"
)

func login(account string, password string) error {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("dial err = ", err)
	}
	var mes message.Message

	mes.Type = message.LoginMesType
	loginMes := message.LoginMes{}
	loginMes.UserId = account
	loginMes.Password = password
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return errors.New("json.Marshal err = ")
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return errors.New("json.Marshal err = ")
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var bytes []byte
	bytes := make([]byte, 4, 4)

	binary.BigEndian.PutUint32(bytes, pkgLen)
	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("发送没有成功 ", err)
		return err

	}
	fmt.Println("发送消息长度成功 ", len(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送消息body失败 ", err)
		return err
	}
	fmt.Println("发送消息body成功 ", data)

	var loginRes message.LoginResMes

	resMsg, err := util.ReadPkg(conn)
	if err != nil {
		fmt.Println("util.ReadPkg ", err)
		return err
	}
	err = json.Unmarshal([]byte(resMsg.Data), &loginRes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return err
	}
	if loginRes.Code == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println("err")

	}
	return nil
}
