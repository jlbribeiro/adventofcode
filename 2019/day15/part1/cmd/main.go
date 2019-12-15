package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day15/part1/oxygen"
)

func main() {
	fmt.Println(oxygen.FindShortestPathToOxygenFromInput(os.Stdin, true))
}
