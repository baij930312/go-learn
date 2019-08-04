package process

import (
	"encoding/json"
	"fmt"
	"net"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/server/model"
	"go_code/learn/file/chat/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcMesLogin(msg message.Message) error {
	data := msg.Data
	fmt.Println(msg)
	var loginMes message.LoginMes
	var resMsg message.Message
	var loginRes message.ResMes
	resMsg.Type = message.ResMesType

	err := json.Unmarshal([]byte(data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	user, err := model.UserDaoInstance.Login(loginMes.UserId, loginMes.Password)
	if err == nil {
		fmt.Println(user)
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
		tf.WritePkg(data)
	} else {
		loginRes.Code = 500
		if err == model.ERROR_USER_NOT_EXEISTS {
			loginRes.Code = 501
			loginRes.Error = "用户不存在"
		}
		if err == model.ERROR_USER_PWD {
			loginRes.Code = 502
			loginRes.Error = "密码不匹配"
		}

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
		tf.WritePkg(data)
	}
	return nil
}

func (this *UserProcess) ServerProcMesRegister(msg message.Message) error {
	data := msg.Data
	fmt.Println(msg)
	var RegisterMes message.RegisterMes
	var resMsg message.Message
	var RegisterMesRes message.ResMes
	resMsg.Type = message.ResMesType

	err := json.Unmarshal([]byte(data), &RegisterMes)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	user, err := model.UserDaoInstance.Register(RegisterMes.UserId, RegisterMes.Password, RegisterMes.UserName)
	if err == nil {
		fmt.Println(user)
		RegisterMesRes.Code = 200
		data, err := json.Marshal(RegisterMesRes)
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
		tf.WritePkg(data)
	} else {
		RegisterMesRes.Code = 500
		if err == model.ERROR_USER_EXEISTS {
			RegisterMesRes.Code = 505
			RegisterMesRes.Error = "用户已存在"
		}
		if err == model.ERROR_USER_PWD {
			RegisterMesRes.Code = 502
			RegisterMesRes.Error = "密码不匹配"
		}

		data, err := json.Marshal(RegisterMesRes)
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
		tf.WritePkg(data)
	}
	return nil
}
