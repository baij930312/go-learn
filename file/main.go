package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 3)
	channel <- 2

	fmt.Println(<-channel)
	fmt.Println(len(channel), "  ", cap(channel))
}
