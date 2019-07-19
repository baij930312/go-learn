package main

import "fmt"

func isPrimeNumber(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func run(source chan int, result chan int, exit chan bool) {
	for {
		n, err := <-source
		if !err {
			break
		}
		if isPrimeNumber(n) {
			result <- n
		}
	}
	exit <- true
}

func main() {
	source := make(chan int, 8000)
	result := make(chan int, 8000)
	exit := make(chan bool, 4)
	go func() {
		for i := 0; i < 8000; i++ {
			source <- i
		}
		close(source)
	}()

	for i := 0; i < 4; i++ {
		go run(source, result, exit)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exit
		}
		close(result) //为啥不写这一行会报错````` //因为主携程在等待result明确推出 否则会造成主携程思索
	}()

	for n := range result {
		
		fmt.Println(n)
	}
	close(exit)
	fmt.Println("over")
}
