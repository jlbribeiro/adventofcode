package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day11/part1/painter"
	"github.com/jlbribeiro/adventofcode/2019/day11/part1/thermal"
)

func main() {
	program := thermal.ProgramFromInput(os.Stdin)
	paint := painter.NewBrainyPainter(program)
	paint.Paint()
	fmt.Println(paint.CountPaintedPanels())
}
