package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jlbribeiro/adventofcode/2017/day5/part1/maze"
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
