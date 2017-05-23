package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello"
	fmt.Println(reflect.TypeOf(str))
	for i := range str {
		fmt.Println(string(str[i]))
	}

	var stdin string
	fmt.Scan(&stdin)
	fmt.Println(stdin)
}
