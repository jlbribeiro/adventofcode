package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day23/part1/cat6"
)

func main() {
	fmt.Println(cat6.FirstPacketTo(os.Stdin, 50, 255))
}
