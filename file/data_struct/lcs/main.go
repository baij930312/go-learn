package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 2, 8, 9, 123, 3}
	b := []int{1, 2, 6, 123, 13, 7, 3}
	fmt.Println(run(a, b))
}

func run(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	lena := len(a)
	lenb := len(b)
	if a[lena-1] == b[lenb-1] {
		arr := run(a[:(lena-1)], b[:(lenb-1)])
		arr = append(arr, a[lena-1])
		return arr
	} else {
		maxa := run(a[:(lena)], b[:(lenb-1)])
		maxb := run(a[:(lena-1)], b[:(lenb)])
		if len(maxa) > len(maxb) {
			return maxa
		} else {
			return maxb
		}
	}
}
