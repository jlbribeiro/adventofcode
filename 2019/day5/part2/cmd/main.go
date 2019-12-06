package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day5/part2/thermal"
)

func main() {
	input := 5
	output := thermal.RunFromInput(os.Stdin, input)
	if len(output) != 1 {
		panic(fmt.Errorf("expecting a single value as output"))
	}
	fmt.Println(output[0])
}
