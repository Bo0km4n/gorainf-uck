package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not mached argument's length")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした")
	}

	interpret(string(data))
}

func interpret(tape string) {
	memory := make([]uint8, 256)
	pointer := 0
	programCounter := 0

	for programCounter = 0; programCounter < len(tape); programCounter++ {

		//test
		//fmt.Println(string(tape[programCounter]))

		switch tape[programCounter] {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case ',':
			var stdin string
			fmt.Scan(&stdin)
			memory[pointer] = uint8(stdin[0])
		case '.':
			fmt.Print(string(memory[pointer]))
		case '[':
			if memory[pointer] == 0 {
				for nest := 1; nest > 0; {
					programCounter++
					c := tape[programCounter]
					if c == '[' {
						nest++
					} else if c == ']' {
						nest--
					}
				}
			}
		case ']':
			for nest := 1; nest > 0; {
				programCounter--
				c := tape[programCounter]
				if c == '[' {
					nest--
				} else if c == ']' {
					nest++
				}
			}
			programCounter--
		}
	}
}

func dumpMemory(memory []uint8) {
	pp.Println(memory)
}
