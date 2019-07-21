package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("conn success.....", conn)
	conn.Do("auth", "baijin")
	_, error := conn.Do("set", "name", "123")
	if error != nil {
		fmt.Println(error)
	}
	str, error := redis.String(conn.Do("get", "name"))
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(str)
}
