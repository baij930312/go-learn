package main

import (
	"fmt"
	"time"
)

func writeData(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		ch <- i
	}
}
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go writeData(ch1)
	go writeData(ch2)

	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		default:
			time.Sleep(time.Second / 2)
			fmt.Println("no data ")
		}
	}

}
