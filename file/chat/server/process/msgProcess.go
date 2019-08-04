package process

import (
	"encoding/json"
	"fmt"
	"net"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/server/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroup(msg message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return err
	}
	data, err := json.Marshal(msg)
	for id, up := range userManagerInstance.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendSmsToEachUser(data, up.Conn)
	}
	return
}

func (this *SmsProcess) SendSmsToEachUser(info []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(info)
	if err != nil {
		fmt.Println("转发消息失败")
	}
}
