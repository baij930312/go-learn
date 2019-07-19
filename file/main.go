package main

import (
	"fmt"
	"reflect"
)

//Person Person
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Test Test
func (p *Person) Test() int {
	return 3
}

//Test1 Test1
func (p Person) Test1() int {
	return 2
}

func test(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println(v.Elem().NumField())
	fmt.Println(v.NumMethod())

	for i := 0; i < v.Elem().NumField(); i++ {
		// v.Elem().Field(i).Tag.Get("json")
		fmt.Println(v.Elem().Type().Field(i).Tag.Get("json"))
	}
	for i := 0; i < v.NumMethod(); i++ {
		// v.Elem().Field(i).Tag.Get("json")
		var param []reflect.Value
		res := v.Method(i).Call(param)
		fmt.Println(v.Method(i))
		fmt.Println(res[0].Int())
	}

}

func main() {
	i := Person{"asdasd", 1}
	test(&i)

	fmt.Println(i)

}
