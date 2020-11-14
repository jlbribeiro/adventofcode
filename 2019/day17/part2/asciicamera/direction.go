package asciicamera

type Direction int

const directions int = 4
const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Left() Direction {
	return Direction((int(d) - 1 + directions) % directions)
}

func (d Direction) Right() Direction {
	return Direction((int(d) + 1) % directions)
}

func (d Direction) Offsets() (int, int) {
	switch d {
	case North:
		return -1, 0
	case South:
		return 1, 0
	case West:
		return 0, -1
	case East:
		return 0, 1
	}

	panic("invalid direction")
}
