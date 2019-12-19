package tractorbeam

import (
	"io"

	"github.com/jlbribeiro/adventofcode/2019/day17/part2/intcode"
)

func AreaFromInput(input io.Reader) int {
	program := intcode.ProgramFromInput(input)

	area := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			cpu := intcode.NewCPU(program)
			output, _ := cpu.Exec([]int64{int64(x), int64(y)})
			area += int(output[0])
		}
	}

	return area
}
