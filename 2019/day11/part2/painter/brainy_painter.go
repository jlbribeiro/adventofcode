package painter

import (
	"io"

	"github.com/jlbribeiro/adventofcode/2019/day11/part2/thermal"
)

type BrainyPainter struct {
	brain   *thermal.CPU
	painter *Painter
}

func NewBrainyPainter(program []int64) *BrainyPainter {
	return &BrainyPainter{
		brain:   thermal.NewCPU(program),
		painter: NewPainter(),
	}
}

func (p *BrainyPainter) Paint() {
	input := make([]int64, 0, 0)

	for {
		scanResult := p.painter.Scan()
		input = append(input, int64(scanResult))

		output, running := p.brain.Exec(input)
		p.painter.Paint(int(output[0]), int(output[1]))

		if !running {
			break
		}
	}
}

func (p *BrainyPainter) CountPaintedPanels() int {
	return p.painter.CountPaintedPanels()
}

func (p *BrainyPainter) Print(out io.Writer) {
	p.painter.Print(out)
}
