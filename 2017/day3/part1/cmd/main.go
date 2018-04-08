package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2017/day3/part1/spiral"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	fmt.Println(spiral.CarryDistance(n))
}
