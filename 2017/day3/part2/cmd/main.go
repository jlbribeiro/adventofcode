package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2017/day3/part2/spiral"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(spiral.GetSumLargerThan(n))
}
