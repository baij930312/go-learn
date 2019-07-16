package main

import (
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (this *Monster) Store() bool {
	content, err := json.Marshal(this)
	if err != nil {
		return false
	}
	err = ioutil.WriteFile("./Monster.txt", content, 0666)
	if err != nil {
		return false
	}
	return true
}

// "./Monster.txt"
func (this *Monster) ReStore(path string) bool {
	content,err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	err = json.Unmarshal(content,this)
	if err != nil {
		return false
	}
	return true
}


func main()  {
	
}