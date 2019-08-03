package process

import (
	"fmt"
	"net"

	"go_code/learn/file/chat/client/utils"
)

func ShowMenu() {
	var key int

	fmt.Println("------------------登陆成功------------------")
	fmt.Println("\t\t 1 显示在线用户列表")
	fmt.Println("\t\t 2 发送信息")
	fmt.Println("\t\t 3 信息列表")
	fmt.Println("\t\t 推出系统 ")

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")

	case 2:
		fmt.Println("发送信息")

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
	}
}
