package main

import (
	"fmt"
	"reflect"
)

func test(a interface{}) {
	aV := reflect.ValueOf(a)
	fmt.Println(aV)
	aV.Elem().SetInt(1111)
	fmt.Println(aV)
}

func main() {
	i := 1
	test(&i)

	fmt.Println(i)

}
