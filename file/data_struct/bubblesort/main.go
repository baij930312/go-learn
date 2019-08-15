package main

import (
	"fmt"
)

func main() {
	arr := []int{23, 2, 1221, 453, 43, 213, 12, 432, 78, 1}
	fmt.Println(arr)
	bubbleSoft(arr)
	fmt.Println(arr)
}

func bubbleSoft(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
