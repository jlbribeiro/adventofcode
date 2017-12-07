package main

import (
	"fmt"
	"io"

	"git.jlbribeiro.com/adventofcode/day6/part2/memory"
)

func main() {
	var n int
	var banks []int

	for {
		match, err := fmt.Scanf("%d", &n)
		if match != 1 || err == io.EOF {
			break
		}

		banks = append(banks, n)
	}

	fmt.Println(banks)
	nIterations := memory.RebalanceRepeatLoop(banks)
	fmt.Println(nIterations)
}
