package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello"
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(str[0])
}
