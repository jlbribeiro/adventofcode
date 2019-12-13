package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day13/part1/breakout"
)

func main() {
	fmt.Println(breakout.BlockTilesCountFromInput(os.Stdin))
}
