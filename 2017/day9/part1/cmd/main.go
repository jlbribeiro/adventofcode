package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day9/part1/runescape"
)

func main() {
	stream := runescape.NewStream(os.Stdin)
	stream.Process()
	fmt.Println(stream.Score())
}
