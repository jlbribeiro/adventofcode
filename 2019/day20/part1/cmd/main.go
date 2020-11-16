package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day20/part1/donutmaze"
)

func main() {
	fmt.Println(donutmaze.MinStepsFromInput(os.Stdin))
}
