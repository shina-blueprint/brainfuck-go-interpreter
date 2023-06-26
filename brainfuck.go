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

		case decrement:

		case right:

		case left:

		case loopStart:

		case loopEnd:

		case output:

		case input:
		}

		codePtr++
	}
}
