package oxygen

import "github.com/jlbribeiro/adventofcode/2019/day15/part2/thermal"

type Walker interface {
	Walk(direction Direction) Status
}

type CPUWalker struct {
	cpu   *thermal.CPU
	input []int64
}

func NewCPUWalker(cpu *thermal.CPU) *CPUWalker {
	return &CPUWalker{
		cpu: cpu,
	}
}

func (c *CPUWalker) Walk(direction Direction) Status {
	c.input = append(c.input, int64(direction))
	output, _ := c.cpu.Exec(c.input)
	return Status(output[0])
}

type FunctionWalker struct {
	x, y   int
	mazeFn func(x, y int) Status
}

func NewFunctionWalker(fn func(x, y int) Status) *FunctionWalker {
	return &FunctionWalker{
		x:      0,
		y:      0,
		mazeFn: fn,
	}
}

func (f *FunctionWalker) Walk(direction Direction) Status {
	dx, dy := direction.Offsets()
	x, y := f.x+dx, f.y+dy

	status := f.mazeFn(x, y)
	if status != HitWall {
		f.x, f.y = x, y
	}

	return status
}
