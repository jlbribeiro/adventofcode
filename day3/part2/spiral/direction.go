package spiral

type Direction int

const (
	Right Direction = iota + 1
	Up
	Left
	Down
)

var Directions = []Direction{
	Right,
	Up,
	Left,
	Down,
}

var directionNames = []string{
	"Right",
	"Up",
	"Left",
	"Down",
}

// RotateLeft returns the direction obtained by rotating to the left.
func (dir Direction) RotateLeft() Direction {
	// dir starts at +1, so int(dir) is effectively +1;
	// % len(Directions) will be between 0 and len(Directions) - 1,
	// so the offset must be added again.
	return Direction(int(dir)%len(Directions) + 1)
}

// RotateRight returns the direction obtained by rotating to the right.
func (dir Direction) RotateRight() Direction {
	// dir starts at +1, so int(dir)-2 is effectively -1;
	// % len(Directions) will be between 0 and len(Directions) - 1,
	// so the offset must be added again.
	return Direction((int(dir)-2+len(Directions))%len(Directions) + 1)
}

// DeltaOffset returns the direction in terms of x and y offsets.
func (dir Direction) DeltaOffset() (int, int) {
	switch dir {
	case Right:
		return +1, 0
	case Up:
		return 0, +1
	case Left:
		return -1, 0
	case Down:
		return 0, -1
	}

	return 0, 0
}

func (dir Direction) String() string {
	if dir < Right || dir > Down {
		return "Unknown"
	}

	return directionNames[dir-1]
}
