package test

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile(`./sample.txt`)
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした")
	}

	memory := make([]int, 256)
	pointer := 0
	program_counter := 0

	tape := string(data)
	tape = strings.Replace(tape, "\n", "", -1)

	fmt.Println(tape)
	fmt.Println("********************")
	fmt.Println(reflect.TypeOf(tape[program_counter]))

	for ; program_counter < len(tape); program_counter++ {
		switch tape[program_counter] {
		case ">":
			pointer += 1
		case "<":
			pointer -= 1
		case "+":
			memory[pointer] += 1
		case "-":
			memory[pointer] -= 1
		}
	}
}
