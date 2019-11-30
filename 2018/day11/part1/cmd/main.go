package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2018/day11/part1/fuel"
)

func main() {
	serialNumber := 1133
	_, x, y := fuel.MaxTotalPower(serialNumber)
	fmt.Printf("%d,%d\n", x, y)
}
