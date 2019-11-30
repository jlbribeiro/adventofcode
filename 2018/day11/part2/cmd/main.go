package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2018/day11/part2/fuel"
)

func main() {
	serialNumber := 1133
	_, x, y, squareSize := fuel.MaxTotalPower(serialNumber)
	fmt.Printf("%d,%d,%d\n", x, y, squareSize)
}
