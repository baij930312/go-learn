package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"go_code/learn/file/chat/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [1024 * 8]byte
}

func (this *Transfer) ReadPkg() (message.Message, error) {
	var msg message.Message
	buf := this.Buf
	length, err := this.Conn.Read(buf[:4])
	if length != 4 || err != nil {
		fmt.Printf("header err err=%v", err)
		return msg, err //!!!
	}
	contentLen := binary.BigEndian.Uint32(buf[:length])
	length, err = this.Conn.Read(buf[:contentLen])
	if err != nil {
		fmt.Printf("body err err=%v", err)
		return msg, err
	}
	err = json.Unmarshal(buf[:length], &msg)
	if err != nil {
		fmt.Printf("Unmarshal err err=%v", err)
		return msg, err
	}
	return msg, nil
}

func (this *Transfer) WritePkg(data []byte) error {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var bytes []byte
	bytes := make([]byte, 4, 4)

	binary.BigEndian.PutUint32(bytes, pkgLen)
	n, err := this.Conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("发送没有成功 ", err)
		return err
	}
	fmt.Println("发送消息长度成功 ", len(data))
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("发送消息body失败 ", err)
		return err
	}
	return nil
}
