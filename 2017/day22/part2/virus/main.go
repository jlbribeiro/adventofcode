package virus

import (
	"bytes"
	"fmt"
	"strings"
)

type AntiVirus struct {
	grid     [][]CellStatus
	pos      *Vector
	dir      Direction
	infected int
}

func NewAVFromInput(input string) *AntiVirus {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]CellStatus, len(lines))
	for i := range grid {
		grid[i] = make([]CellStatus, len(lines[i]))
		for j := range grid[i] {
			if lines[i][j] == '#' {
				grid[i][j] = INFECTED
			} else {
				grid[i][j] = CLEAN
			}
		}
	}

	y := len(lines) / 2
	x := len(lines[0]) / 2

	return &AntiVirus{
		grid:     grid,
		pos:      &Vector{y: y, x: x},
		dir:      NORTH,
		infected: 0,
	}
}

func (av *AntiVirus) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("[%vx%v grid] Carrier is at %s facing %s.\n", len(av.grid), len(av.grid[0]), av.pos, av.dir))
	for y, row := range av.grid {
		for x, cell := range row {
			carrierCell := y == av.pos.y && x == av.pos.x

			if carrierCell {
				buffer.WriteRune('[')
			} else {
				buffer.WriteRune(' ')
			}

			buffer.WriteRune(rune(cell.String()[0]))

			if carrierCell {
				buffer.WriteRune(']')
			} else {
				buffer.WriteRune(' ')
			}
		}

		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (av *AntiVirus) Run(bursts int) {
	for i := 0; i < bursts; i++ {
		av.step()
	}
}

func (av *AntiVirus) rescale() {
	if av.pos.y >= 0 && av.pos.y < len(av.grid) && av.pos.x >= 0 && av.pos.x < len(av.grid[av.pos.y]) {
		return
	}

	// A rescale is needed; add yOff * 2 rows and xOff * 2 cols.
	yOff := 2
	xOff := 2

	newGrid := make([][]CellStatus, len(av.grid)+yOff*2)
	for y := range newGrid {
		newGrid[y] = make([]CellStatus, len(av.grid[0])+xOff*2)
		for x := range newGrid[y] {
			newGrid[y][x] = CLEAN
		}
	}

	for y := range av.grid {
		for x := range av.grid[y] {
			newGrid[yOff+y][xOff+x] = av.grid[y][x]
		}
	}

	av.grid = newGrid
	av.pos = av.pos.Add(&Vector{y: yOff, x: xOff})
}

func (av *AntiVirus) cellStatus() CellStatus {
	return av.grid[av.pos.y][av.pos.x]
}

func (av *AntiVirus) nextCellStatus() {
	nextCellStatus := av.cellStatus().Next()

	if nextCellStatus == INFECTED {
		av.infected++
	}

	av.grid[av.pos.y][av.pos.x] = nextCellStatus
}

func (av *AntiVirus) moveForward() {
	av.pos = av.pos.Add(av.dir.Vector())
}

func (av *AntiVirus) step() {
	av.rescale()

	cellStatus := av.cellStatus()

	switch cellStatus {
	case CLEAN:
		av.dir = av.dir.Left()
	case INFECTED:
		av.dir = av.dir.Right()
	case FLAGGED:
		av.dir = av.dir.Opposite()
	}

	av.nextCellStatus()

	av.moveForward()
}

func (av *AntiVirus) NInfections() int {
	return av.infected
}
