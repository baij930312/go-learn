package process

import (
	"fmt"
	"io"
	"net"

	"go_code/learn/file/chat/common/message"
	"go_code/learn/file/chat/server/utils"
)

type Processer struct {
	Conn net.Conn
}

func (this *Processer) Run() error {
	for {
		tf := utils.Transfer{
			Conn: this.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Printf("read pkg  err=%v", err)
			if err == io.EOF {
				fmt.Println("客户端关闭", err)
			}
			return err
		}
		err = this.ServerProcMes(msg)
		if err != nil {
			return err
		}
	}

}
func (this *Processer) ServerProcMes(msg message.Message) error {
	switch msg.Type {
	case message.LoginMesType:
		userProcess := &UserProcess{
			Conn: this.Conn,
		}
		err := userProcess.ServerProcMesLogin(msg)
		if err != nil {
			return err
		}
	case message.LoginResMesType:

	}
	return nil
}
