package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Node struct {
	col   int
	row   int
	value int
}

func main() {
	var rowArray [11][11]int
	rowArray[1][2] = 1
	rowArray[2][3] = 2

	for _, v := range rowArray {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}

	sparseArr := make([]Node, 0)
	node := Node{11, 11, 0}
	sparseArr = append(sparseArr, node)

	for i, v := range rowArray {
		for j, v2 := range v {
			if v2 != 0 {
				node := Node{i, j, v2}
				sparseArr = append(sparseArr, node)
			}

		}
	}
	fmt.Println(sparseArr)
	f, err := os.OpenFile("./data.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println(err)

		return
	}
	// r := bufio.NewWriter(f)
	// for _, v := range sparseArr {
	// 	s := fmt.Sprintf("%d %d %d \n", v.col, v.row, v.value)
	// 	len, err := r.WriteString(s)
	// 	if err != nil {
	// 		fmt.Println(err)

	// 	}
	// 	fmt.Println(len)
	// }
	// r.Flush()

	r := bufio.NewReader(f)
	var recoverArray [11][11]int
	var (
		col   = 0
		row   = 0
		value = 0
	)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(" err = ", err)
		}
		fmt.Sscanf(str, "%d %d %d \n", &col, &row, &value)
		if value == 0 {
			continue
		}
		recoverArray[col][row] = value
	}

	for _, v := range recoverArray {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}

	fmt.Println("over")
}
