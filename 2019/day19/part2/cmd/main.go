package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day19/part2/tractorbeam"
)

func main() {
	fmt.Println(tractorbeam.FirstSquareFromInput(os.Stdin, 100, 7))
}
