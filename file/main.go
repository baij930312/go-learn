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
	rowJson := `{"name":"123123123213","age":12,"gender":"12"}`
	m := Monster{}
	err := json.Unmarshal([]byte(rowJson),&m)
	if err == nil {
		fmt.Println(m)
	}else{
		fmt.Println(err)
	}
}
