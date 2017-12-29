package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2017/day15/part2/generators"
)

func main() {
	a := generators.NewGenerator(783, 16807, 4)
	b := generators.NewGenerator(325, 48271, 8)

	fmt.Println(generators.Match(a, b))
}
