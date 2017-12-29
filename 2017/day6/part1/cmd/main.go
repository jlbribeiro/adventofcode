package main

import (
	"fmt"
	"io"

	"git.jlbribeiro.com/adventofcode/2017/day6/part1/memory"
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
	nIterations := memory.RebalanceLoop(banks)
	fmt.Println(nIterations)
}
