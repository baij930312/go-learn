package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"go_code/learn/file/chat/common/message"
)

func ReadPkg(conn net.Conn) (message.Message, error) {
	var msg message.Message
	buf := make([]byte, 1024*8)
	length, err := conn.Read(buf[:4])
	if length != 4 || err != nil {
		fmt.Printf("header err err=%v", err)
		return msg, err //!!!
	}
	contentLen := binary.BigEndian.Uint32(buf[:length])
	length, err = conn.Read(buf[:contentLen])
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
func WritePkg(conn net.Conn, data []byte) error {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var bytes []byte
	bytes := make([]byte, 4, 4)

	binary.BigEndian.PutUint32(bytes, pkgLen)
	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("发送没有成功 ", err)
		return err
	}
	fmt.Println("发送消息长度成功 ", len(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送消息body失败 ", err)
		return err
	}
	return nil
}
