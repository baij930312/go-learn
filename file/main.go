package main

import (
	"fmt"
	"time"
)

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("asdsa")
	// var m map[string]int
	// m["123"] = 1
}

func main() {
	go test()

	time.Sleep(time.Second)

}
