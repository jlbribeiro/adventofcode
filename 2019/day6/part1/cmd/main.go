package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day6/part1/orbitmap"
)

func main() {
	objects := orbitmap.ObjectsFromInput(os.Stdin)
	fmt.Println(objects.CountOrbits())
}
