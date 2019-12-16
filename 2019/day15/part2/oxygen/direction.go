package oxygen

type Direction int

const (
	North Direction = 1
	South           = 2
	West            = 3
	East            = 4
)

func (d Direction) Offsets() (int, int) {
	switch d {
	case North:
		return 0, -1

	case South:
		return 0, 1

	case West:
		return -1, 0

	case East:
		return 1, 0

	default:
		panic("invalid Direction")
	}
}

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South

	case South:
		return North

	case West:
		return East

	case East:
		return West

	default:
		panic("invalid Direction")
	}
}
