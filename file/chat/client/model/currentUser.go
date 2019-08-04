package model

import (
	"net"

	"go_code/learn/file/chat/common/message"
)

type CurrentUser struct {
	Conn net.Conn
	message.User
}
