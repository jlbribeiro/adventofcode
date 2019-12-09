package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day9/part2/thermal"
)

func main() {
	program := thermal.ProgramFromInput(os.Stdin)
	fmt.Println(thermal.NewCPU(program).Exec([]int64{2}))
}
