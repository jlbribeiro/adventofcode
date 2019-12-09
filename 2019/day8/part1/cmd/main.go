package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day8/part1/spaceimage"
)

func main() {
	fmt.Println(spaceimage.ChecksumFromInput(os.Stdin, 25, 6))
}
