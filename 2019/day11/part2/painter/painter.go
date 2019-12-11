package painter

import "io"

import "bufio"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Direction int

const (
	Up Direction = 1 << iota
	Left
	Down
	Right
)

func (d Direction) Left() Direction {
	newD := d << 1
	if newD > Right {
		newD = Up
	}

	return newD
}

func (d Direction) Right() Direction {
	newD := d >> 1
	if newD < Up {
		newD = Right
	}

	return newD
}

func (d Direction) Offsets() (int, int) {
	switch d {
	case Up:
		return 0, -1

	case Left:
		return -1, 0

	case Down:
		return 0, 1

	case Right:
		return 1, 0
	}

	return 0, 0
}

type Painter struct {
	X, Y      int
	Direction Direction
	Hull      map[[2]int]int
	Painted   map[[2]int]struct{}
}

func NewPainter() *Painter {
	hull := make(map[[2]int]int)
	hull[[2]int{0, 0}] = 1 // white

	return &Painter{
		X:         0,
		Y:         0,
		Direction: Up,
		Hull:      hull,
		Painted:   make(map[[2]int]struct{}),
	}
}

func (p *Painter) Scan() int {
	return p.Hull[[2]int{p.X, p.Y}]
}

func (p *Painter) Paint(color int, directionChange int) {
	coords := [2]int{p.X, p.Y}

	p.Hull[coords] = color
	p.Painted[coords] = struct{}{}

	if directionChange == 0 {
		p.Direction = p.Direction.Left()
	} else {
		p.Direction = p.Direction.Right()
	}

	offX, offY := p.Direction.Offsets()
	p.X += offX
	p.Y += offY
}

func (p *Painter) CountPaintedPanels() int {
	return len(p.Painted)
}

func (p *Painter) Print(out io.Writer) {
	w := bufio.NewWriter(out)

	minX, maxX := 0, 0
	minY, maxY := 0, 0

	for coords := range p.Hull {
		x, y := coords[0], coords[1]
		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(maxY, y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			coords := [2]int{x, y}
			if p.Hull[coords] == 0 {
				w.WriteRune('.')
			} else {
				w.WriteRune('#')
			}
		}
		w.WriteRune('\n')
	}

	w.Flush()
}
