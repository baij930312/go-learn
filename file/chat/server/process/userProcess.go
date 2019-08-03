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
