package main

import (
	"fmt"

	"git.jlbribeiro.com/adventofcode/day3/part1/spiral"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	fmt.Println(spiral.CarryDistance(n))
}
