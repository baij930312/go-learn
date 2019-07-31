package process

import (
	"encoding/json"
	"fmt"
	"net"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcMesLogin(msg message.Message) error {
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
	tf := &utils.Transfer{
		Conn: this.Conn,
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
		tf.WritePkg(data)
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
		tf.WritePkg(data)
	}
	return nil
}
