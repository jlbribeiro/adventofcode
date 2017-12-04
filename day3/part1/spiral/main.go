package spiral

import (
	"math"
)

// ManhattanDistance returns the taxicab metric
// between (x1, y1) and (x2, y2).
// https://en.wikipedia.org/wiki/Taxicab_geometry
func ManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

// CarryDistance calculates the Manhattan distance from n to 1, assuming
// a spiral setup.
// Eg.:
//   17  16  15  14  13
//   18   5   4   3  12
//   19   6   1   2  11
//   20   7   8   9  10
//   21  22  23---> ...
// Data from square 12 is carried 3 steps, such as: down, left, left.
func CarryDistance(n int) int {
	if n == 1 {
		return 0
	}

	squareRoot := math.Sqrt(float64(n))

	drSquareRoot := int(math.Ceil(squareRoot)) | 0x1
	side := drSquareRoot
	ulSquareRoot := drSquareRoot - 1

	drSquare := drSquareRoot * drSquareRoot
	ulSquare := ulSquareRoot*ulSquareRoot + 1

	// Coordinates of 1, the "access port".
	x1 := drSquareRoot / 2
	y1 := x1

	var x, y int

	if n > drSquare-side {
		x = drSquare - n
		y = side - 1
	} else if n > ulSquare {
		x = 0
		y = n - drSquare
	} else if n > ulSquare-side {
		x = ulSquare - n
		y = 0
	} else {
		x = side - 1
		y = ulSquare - side + 1 - n
	}

	return ManhattanDistance(x, y, x1, y1)
}
