package process

import (
	"fmt"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/client/model"
)

var OnlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurrentUser model.CurrentUser

func OutPutOnlineUsers() {
	fmt.Println("当前在线的user")
	for id, _ := range OnlineUsers {
		fmt.Println(id)
	}
}

func updateUserStatus(notifyMsg *message.NotifyUserStatusMes) {
	user := &message.User{
		UserId: notifyMsg.UserId,
		Status: notifyMsg.Status,
	}
	OnlineUsers[notifyMsg.UserId] = user
	OutPutOnlineUsers()
}
