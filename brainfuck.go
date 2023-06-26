package main

import (
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

		case loopEnd:

		case output:

		case input:
		}

		codePtr++
	}
}
