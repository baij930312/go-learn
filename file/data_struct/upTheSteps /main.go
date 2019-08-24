package main

import "fmt"

func main() {
	fmt.Println(run(7))
}

func run(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return run(n-1) + run(n-2)
}
