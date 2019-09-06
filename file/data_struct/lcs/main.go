package main

import "fmt"

func main() {
	a := []int{'a', 'd', '4', '2', '8', '9', '3'}
	b := []int{'a', '3', 'd', 'f', 'g', 'h', '3'}
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
