package main

import (
	"fmt"
	"os"
	"bufio"
)

func main (){
	file ,err := os.OpenFile("./temp.txt", os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	wr := bufio.NewWriter(file);
	wr.WriteString("123123")
	//刷新缓冲区
	wr.Flush()
}