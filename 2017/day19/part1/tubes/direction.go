package tubes

type Direction int

const (
	NORTH Direction = iota + 1
	WEST
	SOUTH
	EAST
)

var directions = []Direction{
	NORTH,
	WEST,
	SOUTH,
	EAST,
}

var directionNames = []string{
	"NORTH",
	"WEST",
	"SOUTH",
	"EAST",
}

func (d Direction) String() string {
	if d < NORTH || d > EAST {
		return "NOPE"
	}

	return directionNames[int(d)-1]
}

func (d Direction) RotateLeft() Direction {
	return directions[(int(d)-2+len(directions))%len(directions)]
}

func (d Direction) RotateRight() Direction {
	return directions[int(d)%len(directions)]
}

func (d Direction) Vector() *Vector {
	switch d {
	case NORTH:
		return &Vector{-1, 0}
	case WEST:
		return &Vector{0, -1}
	case SOUTH:
		return &Vector{+1, 0}
	case EAST:
		return &Vector{0, +1}
	}

	return nil
}
