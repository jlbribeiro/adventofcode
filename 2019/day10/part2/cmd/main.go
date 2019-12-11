package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day10/part2/station"
)

func main() {
	grid := station.AsteroidsGridFromInput(os.Stdin)
	_, bestX, bestY := station.BestLoSLocation(grid)
	x, y := station.LaserSweep(grid, bestX, bestY, 200)
	fmt.Println(x*100 + y)
}
