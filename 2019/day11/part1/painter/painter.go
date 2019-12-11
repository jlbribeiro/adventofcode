package painter

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
	X             int
	Y             int
	Direction     Direction
	Hull          map[[2]int]int
	PaintedPanels int
}

func NewPainter() *Painter {
	return &Painter{
		X:             0,
		Y:             0,
		Direction:     Up,
		Hull:          make(map[[2]int]int),
		PaintedPanels: 0,
	}
}

func (p *Painter) Scan() int {
	return p.Hull[[2]int{p.X, p.Y}]
}

func (p *Painter) Paint(color int, directionChange int) {
	coords := [2]int{p.X, p.Y}
	if _, ok := p.Hull[coords]; !ok {
		p.PaintedPanels++
	}

	p.Hull[coords] = color

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
	return p.PaintedPanels
}
