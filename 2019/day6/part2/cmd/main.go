package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day6/part2/orbitmap"
)

func main() {
	objects := orbitmap.ObjectsFromInput(os.Stdin)
	fmt.Println(objects.OrbitalTransfersToSanta())
}
