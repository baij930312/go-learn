package main

import "fmt"

func main() {
	hailston(10)
}

func hailston(n int) int {
	res := 1
	if n > 1 {
		if (n % 2) == 0 {
			res = n / 2
		} else {
			res = 3*n + 1
		}
	} else {
		return 1
	}
	fmt.Println(res)
	return hailston(res)
}
