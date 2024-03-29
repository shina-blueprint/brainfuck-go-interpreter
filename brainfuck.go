package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	increment = '+'
	decrement = '-'
	right     = '>'
	left      = '<'
	loopStart = '['
	loopEnd   = ']'
	output    = '.'
	input     = ','
)

const memorySize = 1024

func main() {
	var (
		memory  [memorySize]byte
		ptr     int
		codePtr int
		loops   []int
	)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: A Brainfuck code is not passed as a command-line argument.")
		fmt.Fprintln(os.Stderr, "Please pass an argument as the form, $ ./brainfuck [FILENAME].")
		os.Exit(1)
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: The file, "+os.Args[1]+", cannot be opened.")
		os.Exit(1)
	}

	codeLen := len(code)

	for codePtr < codeLen {
		switch code[codePtr] {
		case increment:
			memory[ptr]++

		case decrement:
			memory[ptr]--

		case right:
			if ptr >= memorySize-1 {
				ptr = 0
			} else {
				ptr++
			}

		case left:
			if ptr <= 0 {
				ptr = memorySize - 1
			} else {
				ptr--
			}

		case loopStart:
			loops = append(loops, codePtr)
			if memory[ptr] == 0 {
				depth := 1
				for depth > 0 {
					codePtr++
					if codePtr >= codeLen {
						fmt.Fprintln(os.Stderr, "Error: Loop end order, "+string(loopEnd)+", is not found.")
						os.Exit(1)
					}
					if code[codePtr] == loopStart {
						depth++
					}
					if code[codePtr] == loopEnd {
						depth--
					}
				}
				loops = loops[:len(loops)-1]
			}

		case loopEnd:
			if len(loops) == 0 {
				fmt.Fprintln(os.Stderr, "Error: Loop start order, "+string(loopStart)+", is not found.")
				os.Exit(1)
			}
			codePtr = loops[len(loops)-1] - 1
			loops = loops[:len(loops)-1]

		case output:
			fmt.Print(string(memory[ptr]))

		case input:
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			memory[ptr] = line[0]
		}

		codePtr++
	}
}
