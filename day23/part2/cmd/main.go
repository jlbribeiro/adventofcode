package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day23/part2/coprocessor"
)

func main() {
	program := []string(nil)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()
		program = append(program, instruction)
	}

	cp := coprocessor.NewCoprocessor()
	cp.SetValue("a", 1)
	cp.Run(program)
	fmt.Println(cp.GetValue("h"))
}
