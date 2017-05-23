package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした")
	}

	memory := make([]uint8, 256)
	pointer := 0
	program_counter := 0

	tape := string(data)
	fmt.Println("********************")
	fmt.Println(reflect.TypeOf(tape[program_counter]))

	for program_counter = 0; program_counter < len(tape); program_counter++ {

		//test
		//fmt.Println(string(tape[program_counter]))

		switch tape[program_counter] {
		case '>':
			pointer += 1
		case '<':
			pointer -= 1
		case '+':
			memory[pointer] += 1
		case '-':
			memory[pointer] -= 1
		case ',':
			var stdin string
			fmt.Print("Please input something : ")
			fmt.Scan(stdin)
			if len(stdin) != 1 {
				fmt.Println("stdio err")
				break
			}
			memory[pointer] = uint8(stdin[0])
		case '.':
			fmt.Print(string(memory[pointer]))
		case '[':
			if memory[pointer] == 0 {
				for nest := 1; nest > 0; {
					program_counter++
					c := tape[program_counter]
					if c == '[' {
						nest++
					} else if c == ']' {
						nest--
					}
				}
			}
		case ']':
			for nest := 1; nest > 0; {
				program_counter -= 1
				c := tape[program_counter]
				if c == '[' {
					nest--
				} else if c == ']' {
					nest++
				}
			}
			program_counter--
		}
	}
}
