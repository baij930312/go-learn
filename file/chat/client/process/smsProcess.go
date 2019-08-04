package process

import (
	"encoding/json"
	"errors"
	"fmt"

	"go_code/learn/file/chat/client/utils"
	"go_code/learn/file/chat/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroup(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.UserId = CurrentUser.UserId
	smsMes.Content = content
	smsMes.Status = CurrentUser.Status

	data, err := json.Marshal(smsMes)
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

	tf := &utils.Transfer{
		Conn: CurrentUser.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *SmsProcess) ReceiveSms(msg message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return err
	}
	fmt.Println(smsMes.UserId, "说:", smsMes.Content)
	return
}
