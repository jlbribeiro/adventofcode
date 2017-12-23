package virus

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
	"North",
	"West",
	"South",
	"East",
}

func (d Direction) Left() Direction {
	return directions[int(d)%len(directions)]
}

func (d Direction) Right() Direction {
	return directions[(int(d)-2+len(directions))%len(directions)]
}

func (d Direction) Opposite() Direction {
	return d.Left().Left()
}

func (d Direction) Vector() *Vector {
	switch d {
	case NORTH:
		return &Vector{y: -1, x: 0}
	case WEST:
		return &Vector{y: 0, x: -1}
	case SOUTH:
		return &Vector{y: +1, x: 0}
	case EAST:
		return &Vector{y: 0, x: +1}
	}

	return &Vector{y: 0, x: 0}
}

func (d Direction) String() string {
	if d < NORTH || d > EAST {
		return "Unknown"
	}

	return directionNames[int(d)-1]
}
