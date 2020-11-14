package doorkey

type Direction int

const directions int = 4

const (
	North Direction = iota
	West
	South
	East
)

func (d Direction) Left() Direction {
	return Direction((int(d) + 1) % directions)
}

func (d Direction) Right() Direction {
	return Direction((int(d) - 1 + directions) % directions)
}

func (d Direction) Opposite() Direction {
	return Direction((int(d) + 2) % directions)
}

func (d Direction) Offsets() (int, int) {
	switch d {
	case North:
		return -1, 0
	case West:
		return 0, -1
	case South:
		return 1, 0
	case East:
		return 0, 1
	default:
		panic("invalid direction")
	}
}
