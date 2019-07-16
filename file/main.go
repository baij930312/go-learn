package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["hahaha"] = "asdasd"
	m1["1111"] = "aaaa"
	slice = append(slice, m1)
	content, err := json.Marshal(slice)
	if err == nil {
		fmt.Println(string(content))
	}else{
		fmt.Println(err)
	}
}
