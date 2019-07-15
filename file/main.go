package main

import (
	"fmt"
	// "io/ioutil"
	"os"
)

func main() {
	exisit, _ := PathExisists("./temp.go")
	fmt.Println(exisit)
}

func PathExisists(path string) (bool, error) {
	_, error := os.Stat(path)
	if error != nil {
		return true, nil
	}

	if os.IsNotExist(error) {
		return false, nil
	}

	return false, error
}
