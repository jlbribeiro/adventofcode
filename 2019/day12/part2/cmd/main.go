package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day12/part2/nbody"
)

func main() {
	fmt.Println(nbody.SimulateUntilPreviousStateFromInput(os.Stdin))
}
