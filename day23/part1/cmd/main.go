package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day23/part1/coprocessor"
)

func main() {
	program := []string(nil)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()
		program = append(program, instruction)
	}

	cp := coprocessor.NewCoprocessor()
	cp.Run(program)
	fmt.Println(cp.MulCalls())
}
