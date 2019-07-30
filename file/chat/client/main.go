package main

import (
	"fmt"
	"os"
)

var account string
var password string

func main() {
	var key int
	var loop = true
	for loop {
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
			fmt.Scanf("%s\n", &account)
			fmt.Println("password:")
			fmt.Scanf("%s\n", &password)
			err := login(account, password)
			if err != nil {
				fmt.Println("登录失败")
			} else {
				fmt.Println("登录成功")
			}
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出")
			loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误请重新输入")
		}
	}

}
