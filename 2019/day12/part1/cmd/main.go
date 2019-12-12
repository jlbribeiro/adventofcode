package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day12/part1/nbody"
)

func main() {
	fmt.Println(nbody.SimulateEnergyFromInput(os.Stdin, 1000))
}
