package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"go_code/learn/file/chat/client/utils"
	"go_code/learn/file/chat/common/message"
)

type UserProcess struct {
}

func (this *UserProcess) Login(account int, password string) error {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("dial err = ", err)
	}
	var mes message.Message
	var loginRes message.ResMes
	tf := &utils.Transfer{
		Conn: conn,
	}
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

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("util.WritePkg ", err)
		return err
	}
	resMsg, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("util.ReadPkg ", err)
		return err
	}
	err = json.Unmarshal([]byte(resMsg.Data), &loginRes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return err
	}
	fmt.Println(loginRes)
	if loginRes.Code == 200 {
		go serverProcessSms(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginRes.Code)
		fmt.Println(loginRes.Error)
	}
	return nil
}

func (this *UserProcess) Register(account int, password string, username string) error {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("dial err = ", err)
	}
	var mes message.Message
	var loginRes message.ResMes
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes.Type = message.RegisterMesType
	RegisterMes := message.RegisterMes{}
	RegisterMes.UserId = account
	RegisterMes.Password = password
	RegisterMes.UserName = username
	data, err := json.Marshal(RegisterMes)
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

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("util.WritePkg ", err)
		return err
	}
	resMsg, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("util.ReadPkg ", err)
		return err
	}
	err = json.Unmarshal([]byte(resMsg.Data), &loginRes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return err
	}
	fmt.Println(loginRes)
	if loginRes.Code == 200 {
		fmt.Println("注册成功")
		go serverProcessSms(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginRes.Code)
		fmt.Println(loginRes.Error)
	}
	return nil
}
