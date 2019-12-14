package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day13/part2/breakout"
)

func main() {
	printGame := os.Getenv("PRINT_GAME") == "true"
	fmt.Println(breakout.GameScoreFromInput(os.Stdin, printGame))
}
