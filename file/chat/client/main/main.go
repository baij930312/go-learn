package main

import (
	"fmt"
	"os"

	"go_code/learn/file/chat/client/process"
)

var account int
var password string
var userName string

func main() {
	var key int
	for {
		fmt.Println("------------------欢迎使用------------------")
		fmt.Println("\t\t 1 登录聊天")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出")
		fmt.Println("\t\t 请输入:[1-3]")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("account:")
			fmt.Scanf("%d\n", &account)
			fmt.Println("password:")
			fmt.Scanf("%s\n", &password)
			// login(account, password)
			up := process.UserProcess{}
			up.Login(account, password)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("account:")
			fmt.Scanf("%d\n", &account)
			fmt.Println("password:")
			fmt.Scanf("%s\n", &password)
			fmt.Println("username:")
			fmt.Scanf("%s\n", &userName)
			up := process.UserProcess{}
			up.Register(account, password, userName)

		case 3:
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("输入有误请重新输入")
		}
	}
}
