package spiral

import (
	"math"
)

type Spiral struct {
	side   int
	matrix [][]int
	x      int
	y      int
	dir    Direction
}

func NewSpiral(side int) *Spiral {
	matrix := make([][]int, side, side)
	for i := 0; i < side; i++ {
		matrix[i] = make([]int, side, side)
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	// Start at the center of the spiral
	x := side / 2
	y := side / 2

	matrix[x][y] = 1

	return &Spiral{
		side:   side,
		matrix: matrix,
		x:      x,
		y:      y,
		dir:    Right,
	}
}

func (m *Spiral) sumNeighbors(x int, y int) int {
	sum := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if x == 0 && y == 0 {
				continue
			}

			_x := x + dx
			_y := y + dy

			sum += m.valueAt(_x, _y)
		}
	}

	return sum
}

// Next calculates the sum of the neighbors of the (x,y) square of the spiral.
func (m *Spiral) Next() int {
	dX, dY := m.dir.DeltaOffset()
	m.x += dX
	m.y += dY

	m.matrix[m.x][m.y] = m.sumNeighbors(m.x, m.y)

	m.updateDirection()

	return m.valueAt(m.x, m.y)
}

func (m *Spiral) updateDirection() {
	nextDir := m.dir.RotateLeft()
	dx, dy := nextDir.DeltaOffset()

	if m.valueAt(m.x+dx, m.y+dy) == 0 {
		m.dir = nextDir
	}
}

func (m *Spiral) insideBounds(x int, y int) bool {
	return x >= 0 && x < m.side && y >= 0 && y < m.side
}

func (m *Spiral) valueAt(x int, y int) int {
	// Consider an infinite matrix of 0s
	if !m.insideBounds(x, y) {
		return 0
	}

	return m.matrix[x][y]
}

// GetSumLargerThan floods a matrix with the sum of the adjacent squares in
// a spiral, in order to calculate the first square that has a sum larger than
// n.
func GetSumLargerThan(value int) int {
	side := int(math.Ceil(math.Sqrt(float64(value))))
	spiral := NewSpiral(side)

	var sum int

	for {
		sum = spiral.Next()
		if sum > value {
			return sum
		}
	}
}
