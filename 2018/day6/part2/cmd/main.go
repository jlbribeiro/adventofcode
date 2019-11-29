package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2018/day6/part2/manhattan"
)

func main() {
	fmt.Println(manhattan.RegionSizeOfMaxDistance(os.Stdin, 10000))
}
