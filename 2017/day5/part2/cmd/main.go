package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"git.jlbribeiro.com/adventofcode/day5/part2/maze"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var mazeSteps []int
	for scanner.Scan() {
		ns := scanner.Text()
		n, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}

		mazeSteps = append(mazeSteps, n)
	}

	nSteps := maze.Walk(mazeSteps)
	fmt.Println(nSteps)
}
