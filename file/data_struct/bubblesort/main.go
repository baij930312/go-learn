package main

import (
	"fmt"
)

func main() {
	arr := []int{23, 2, 1221, 453, 43, 213, 12, 432, 78, 1}
	fmt.Println(arr)
	arr = bubbleSoft(arr)
	fmt.Println(arr)
}

func bubbleSoft(rowarr []int) []int {
	arr := make([]int, 10)
	copy(arr, rowarr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
