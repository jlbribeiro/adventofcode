package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day10/part1/station"
)

func main() {
	grid := station.AsteroidsGridFromInput(os.Stdin)
	bestLoS, bestX, bestY := station.BestLoSLocation(grid)
	fmt.Printf("(%d,%d) has a LoS of %d asteroids.\n", bestX, bestY, bestLoS)
}
