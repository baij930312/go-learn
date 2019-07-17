package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]int)
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for index := 1; index <= n; index++ {
		res *= index
	}
	lock.Lock()
	m[n] = res
	lock.Unlock()
}

func main() {
	for index := 0; index < 100; index++ {
		go test(index)
	}

	time.Sleep(time.Second * 3)
	lock.Lock()
	fmt.Println(m )
	lock.Unlock()
}
