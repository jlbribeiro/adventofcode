package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day11/part1/hexed"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	walker := hexed.NewHexWalker()
	walker.WalkFromInput(input)
	fmt.Println(walker.MinStepsToStart())
}
