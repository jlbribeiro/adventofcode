package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day8/part2/spaceimage"
)

func main() {
	img := spaceimage.RenderFromInput(os.Stdin, 25, 6)
	for _, row := range img {
		for _, col := range row {
			if col == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
