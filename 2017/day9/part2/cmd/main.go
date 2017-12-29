package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day9/part2/runescape"
)

func main() {
	stream := runescape.NewStream(os.Stdin)
	stream.Process()
	fmt.Println(stream.RemovedGarbageCount())
}
