package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2))
	fmt.Println(reflect.TypeOf(2.4))
	fmt.Println(reflect.TypeOf("Hello World"))
	fmt.Println(reflect.TypeOf(true))
}
