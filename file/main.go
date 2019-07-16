package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	m := Monster{
		Name:"123213",
		Age:12,
		Gender:"12",
	}
	content, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(content))
	}else{
		fmt.Println(err)
	}
}
