package process

import (
	"encoding/json"
	"fmt"
	"net"

	"go_code/learn/file/chat/client/utils"
	"go_code/learn/file/chat/common/message"
)

func ShowMenu() {
	var key int
	var content string

	fmt.Println("------------------登陆成功------------------")
	fmt.Println("\t\t 1 显示在线用户列表")
	fmt.Println("\t\t 2 发送信息")
	fmt.Println("\t\t 3 信息列表")
	fmt.Println("\t\t 推出系统 ")
	sms := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		OutPutOnlineUsers()

	case 2:
		fmt.Println("请讲")
		fmt.Scanf("%s\n", &content)
		sms.SendGroup(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("推出")

	default:
		fmt.Println("输入有误请重新输入")
	}
}

func serverProcessSms(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("read error ")
		}
		fmt.Println(msg)

		switch msg.Type {
		case message.NotifyUserStatusMesType:
			var notifMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(msg.Data), &notifMes)
			if err != nil {
				fmt.Println(" json.Unmarshal")
			}
			updateUserStatus(&notifMes)
		case message.SmsMesType:
			smsProcess := &SmsProcess{}
			smsProcess.ReceiveSms(msg)
		default:
			fmt.Println("收到消息")
		}
	}
}
