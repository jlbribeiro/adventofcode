package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day24/part2/bugs"
)

func main() {
	fmt.Println(bugs.BugsAfterNGenerationsFromInput(os.Stdin, 200))
	// fmt.Println(bugs.BugsAfterNGenerationsFromInput(os.Stdin, 10))
}
