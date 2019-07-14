package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := ioutil.ReadFile("./temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(fileContent))
}
