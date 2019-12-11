package main

import (
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day11/part2/painter"
	"github.com/jlbribeiro/adventofcode/2019/day11/part2/thermal"
)

func main() {
	program := thermal.ProgramFromInput(os.Stdin)
	paint := painter.NewBrainyPainter(program)
	paint.Paint()
	paint.Print(os.Stdout)
}
