package main

import (
	"fmt"
	"reflect"
)

//Person Person
type Person struct {
	Name string
	Age  int
}

func main() {
	i := 1
	refI := reflect.ValueOf(i)
	reftype := reflect.TypeOf(i)

	fmt.Println(refI.Kind())
	fmt.Println(reftype)
	fmt.Println(refI.Int())

	var p = Person{"asd", 1}
	iV := reflect.ValueOf(p)
	fmt.Println(iV.Kind()) //struct
	fmt.Println(iV.Type()) //main.Person

}
